# Server builder
FROM golang:1.21-alpine as server_builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./ent ./ent
COPY ./pkg ./pkg
COPY ./vendor ./vendor

RUN go build -v -o ./bin/server ./cmd/server/server.go

# Runtime
FROM alpine:3.18

WORKDIR /

COPY --from=server_builder /app/bin/server .

COPY ./data ./data

COPY ./.env* .

EXPOSE 8080

CMD ["/server"]
