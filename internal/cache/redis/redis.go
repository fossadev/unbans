package redis

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/fossadev/unbans/internal/cache"
	"github.com/fossadev/unbans/internal/config"
	"github.com/mediocregopher/radix/v3"
	"github.com/pkg/errors"
)

type redis struct {
	client *radix.Pool
}

func New(cfg *config.RedisConfig) (cache.Cache, error) {
	pool, err := radix.NewPool(cfg.Network, cfg.Host, 10)
	if err != nil {
		return nil, errors.Wrap(err, "failed to make redis pool")
	}

	return &redis{
		client: pool,
	}, nil
}

func (c *redis) AttemptStateToken(stateToken, payload string, stateTokenExpiry time.Duration) (bool, error) {
	var rcv string
	dur := strconv.Itoa(int(stateTokenExpiry / time.Second))
	cmd := radix.Cmd(&rcv, "SET", c.stateTokenKey(stateToken), payload, "EX", dur, "NX")
	if err := c.client.Do(cmd); err != nil {
		return true, errors.Wrap(err, "failed to attempt SET for state token")
	}
	return rcv != "OK", nil
}

func (c *redis) GetStatePayload(state string, out interface{}) error {
	var rcv string
	stateTokenKey := c.stateTokenKey(state)
	if err := c.client.Do(radix.Cmd(&rcv, "GET", stateTokenKey)); err != nil {
		return err
	}

	if rcv == "" {
		return cache.ErrKeyNotFound
	}

	go func() {
		_ = c.client.Do(radix.Cmd(nil, "DEL", stateTokenKey))
	}()

	if err := json.Unmarshal([]byte(rcv), out); err != nil {
		return err
	}

	return nil
}

func (c *redis) stateTokenKey(stateToken string) string {
	return "fossadev-unbans::oauth-state-tokens::" + stateToken
}
