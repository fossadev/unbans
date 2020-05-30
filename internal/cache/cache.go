package cache

import (
	"errors"
	"time"
)

var ErrKeyNotFound = errors.New("no key found")

type Cache interface {
	// OAuth States
	AttemptStateToken(stateToken, payload string, stateTokenExpiry time.Duration) (exists bool, err error)
	GetStatePayload(state string, out interface{}) (err error)
}
