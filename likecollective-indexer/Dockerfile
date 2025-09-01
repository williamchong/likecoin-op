FROM golang:1.23.5-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache make
RUN make likecollective-indexer-api

FROM ubuntu:noble
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    make \
    && rm -rf /var/lib/apt/lists/*
RUN update-ca-certificates

# Cmds
COPY --from=builder /src/Makefile /app

# App
COPY --from=builder /src/likecollective-indexer-api /usr/local/bin/

# DB Mock data
COPY --from=builder /src/dbmockdata /app/dbmockdata

CMD ["likecollective-indexer-api"]
