#!/bin/sh

if [ ! -d scripts ]; then
	echo "This must be executed from the project root :: ./scripts/build.sh"
	exit 1
fi

if ! [ -f ../.env ]; then
	echo "[WARN] no .env file found. Executing without .env file"
	$"$@"
	exit
fi

env $(cat .env | xargs) "$@"
