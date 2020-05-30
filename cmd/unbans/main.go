package main

import (
	"context"
	nativeLog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fossadev/unbans/internal/cache/redis"
	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/db/postgres"
	"github.com/fossadev/unbans/internal/features"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/fossadev/unbans/internal/server"
	"github.com/fossadev/unbans/internal/twitchapi"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	log, err := logger.New(cfg.Environment)
	if err != nil {
		nativeLog.Fatalln("Failed to init logger: " + err.Error())
	}

	pgdb := postgres.New(cfg.Postgres)
	cacheImpl, err := redis.New(cfg.Redis)
	if err != nil {
		log.Fatal("failed to connect to redis", zap.Error(err))
	}

	dbImpl := pgdb.DB()
	featuresImpl := features.New(cacheImpl, cfg, dbImpl, log)
	twitchAPIImpl := twitchapi.New(cfg.Twitch.ClientID)

	serverConf := &server.Config{
		Cache:     cacheImpl,
		Config:    cfg,
		DB:        dbImpl,
		Features:  featuresImpl,
		Log:       log,
		TwitchAPI: twitchAPIImpl,
	}

	srv := server.New(serverConf)
	srv.Start()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err = srv.Stop(ctx); err != nil {
		log.Error("failed to stop server", err)
	}

	log.Info("bye!")
}
