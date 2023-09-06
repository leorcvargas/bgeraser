# Server builder
FROM golang:1.21-alpine as server_builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY ./cmd ./cmd
COPY ./internal ./internal
COPY ./ent ./ent
COPY ./pkg ./pkg

RUN go build -v -o ./bin/server ./cmd/server.go

# Runtime
FROM alpine:3.14.10

WORKDIR /

EXPOSE 8080

COPY --from=server_builder /app/bin/server .

CMD ["/server"]
