package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fossadev/unbans/internal/config"
	"github.com/go-pg/migrations/v7"
	"github.com/go-pg/pg/v9"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  go run *.go <command> [args]
`

func main() {
	flag.Usage = usage
	flag.Parse()
	log.Println("Running migrations...")
	cfg := config.New()

	db := pg.Connect(&pg.Options{
		Network:  cfg.Postgres.Network,
		Addr:     cfg.Postgres.Host,
		User:     cfg.Postgres.Username,
		Password: cfg.Postgres.Password,
		Database: cfg.Postgres.Database,
	})

	oldVersion, newVersion, err := migrations.Run(db, flag.Args()...)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}
