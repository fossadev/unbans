lint:
	golangci-lint run

dev:
	./scripts/with-env.sh go run cmd/unbans/*.go

migrations:
	./scripts/with-env.sh go run cmd/migrations/*.go init
	./scripts/with-env.sh go run cmd/migrations/*.go up

migrations_down:
	./scripts/with-env.sh go run cmd/migrations/*.go reset

test_api:
	go run -v ./api/...
	go run -v -test.race ./api/...

test_internal:
	go run -v ./internal/...
	go run -v -test.race ./internal/...

test_src:
	yarn test
