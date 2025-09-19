FROM golang:1.23.5-alpine AS builder
RUN apk add curl
RUN curl https://atlasgo.sh -o atlasgo.sh
RUN sh atlasgo.sh -y
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache make
RUN make likecollective-indexer-api
RUN make likecollective-indexer-worker-cli

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
COPY --from=builder /src/likecollective-indexer-worker-cli /usr/local/bin/

# Migration
COPY --from=builder /usr/local/bin/atlas /usr/local/bin/
COPY --from=builder /src/ent/migrate/migrations /app/ent/migrate/migrations

CMD ["likecollective-indexer-api"]
