package main

import (
	"context"
	nativeLog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fossadev/unbans/internal/config"
	"github.com/fossadev/unbans/internal/logger"
	"github.com/fossadev/unbans/internal/server"
)

func main() {
	cfg := config.New()
	log, err := logger.New(cfg.Environment)
	if err != nil {
		nativeLog.Fatalln("Failed to init logger: " + err.Error())
	}

	srv := server.New(cfg, log)
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
