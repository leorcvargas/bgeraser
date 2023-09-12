dev:
	go run ./cmd/server.go

build: clean deps
	go build -v -o ./bin/bgeraser ./cmd/server.go

deps:
	go mod tidy -v

clean:
	rm -rf ./bin

test:
	go test -coverprofile=coverage.out -v ./...

dbuild:
	docker build -t leorcvargas/bgeraser .

dcdown:
	docker compose down --remove-orphans

dcup: dcdown
	docker compose up --build

dcupd: dcdown
	docker compose up --build -d

