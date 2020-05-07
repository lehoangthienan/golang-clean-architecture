init:
	cat .env.example > .env
	cat .env_migrator.yaml.example > .env_migrator.yaml
	go mod init github.com/lehoangthienan/marvel-heroes-backend
	go mod tidy
	docker-compose -f docker-compose-local.yaml up -d

bin:
	go build -o bin/migrator ./cmd/migrator
	bin/migrator up

dev: bin
	go mod tidy
	ENV=local go run cmd/server/main.go

test:
	go test ./...

clean:
	docker-compose -f docker-compose-local.yaml down

setup-db:
	docker-compose -f docker-compose-local.yaml up -d

secure-grpc: bin
	go mod tidy
	ENV=secure-grpc go run cmd/server/main.go
