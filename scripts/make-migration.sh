#!/bin/sh

script_path="$(dirname "$0")"
migrations_directory="$script_path/../cmd/migrations"
file_name="$migrations_directory/$(date +%s)_$1.go"

cat > $file_name <<EOF
package main

import "github.com/go-pg/migrations/v7"

func init() {
	migrations.MustRegister(func(db migrations.DB) error {
		_, err := db.Exec(\`

		\`)
		return err
	}, func(db migrations.DB) error {
		_, err := db.Exec(\`

		\`)
		return err
	})
}

EOF
