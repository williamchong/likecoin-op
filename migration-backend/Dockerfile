FROM golang:1.22.10-alpine AS stage1
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN apk add --no-cache make
RUN make build
RUN make build-worker-cli
RUN go install github.com/rubenv/sql-migrate/...@v1.7.1

FROM ubuntu:noble
WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    tzdata \
    media-types \
    && rm -rf /var/lib/apt/lists/*
RUN update-ca-certificates

# App
COPY --from=stage1 /src/like-migration-backend /usr/local/bin/
COPY --from=stage1 /src/like-migration-backend-worker-cli /usr/local/bin/

# Migration
COPY --from=stage1 /go/bin/sql-migrate /usr/local/bin/
COPY --from=stage1 /src/migrations /app/migrations
COPY --from=stage1 /src/dbconfig.yml /app/dbconfig.yml

EXPOSE 8091
CMD ["like-migration-backend"]
