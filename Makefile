dev:
	go run ./cmd/server.go

build: clean deps
	go build -v -o ./bin/bgeraser ./cmd/server.go

deps:
	go mod tidy -v

clean:
	rm -rf ./bin

docker:
	docker build -t leorcvargas/bgeraser .

docker-down:
	docker compose down -v --remove-orphans

docker-dev: docker-down
	docker compose -f docker-compose.dev.yml up --build

docker-local: docker-down
	docker compose -f docker-compose.local.yml up --build -d