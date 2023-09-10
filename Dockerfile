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
FROM python:3.11-slim

WORKDIR /

COPY --from=server_builder /app/bin/server .

COPY ./data ./data

COPY ./.env* .

RUN pip3 install --no-cache-dir rembg[cli]

RUN python -c 'from rembg.bg import download_models; download_models()'

EXPOSE 8080

CMD ["/bin/sh", "-c", "/server"]
