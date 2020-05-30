lint:
	golangci-lint run

dev:
	./scripts/with-env.sh go run cmd/unbans/*.go

migrations:
	./scripts/with-env.sh go run cmd/migrations/*.go up

migrations_init:
	./scripts/with-env.sh go run cmd/migrations/*.go init

migrations_down:
	./scripts/with-env.sh go run cmd/migrations/*.go reset

test: test_api test_internal

test_api:
	go test -v ./api/...
	go test -v -test.race ./api/...

test_internal:
	go test -v ./internal/...
	go test -v -test.race ./internal/...

test_src:
	yarn test
