FROM golang:1.23.0-alpine AS stage1
RUN apk add curl
RUN curl https://atlasgo.sh -o atlasgo.sh
RUN sh atlasgo.sh -y
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache make
RUN make build-api
RUN make build-worker-cli

FROM ubuntu:noble
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    make \
    && rm -rf /var/lib/apt/lists/*
RUN update-ca-certificates

# Cmds
COPY --from=stage1 /src/Makefile /app

# App
COPY --from=stage1 /src/likenft-indexer-api /usr/local/bin/
COPY --from=stage1 /src/likenft-indexer-worker-cli /usr/local/bin/

# Migration
COPY --from=stage1 /usr/local/bin/atlas /usr/local/bin/
COPY --from=stage1 /src/ent/migrate/migrations /app/ent/migrate/migrations

CMD ["likenft-indexer-api"]
