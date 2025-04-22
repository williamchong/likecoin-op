FROM --platform=$BUILDPLATFORM node:23.7-slim as builder

RUN apt-get update \
    && apt-get install -y python3 make g++\
    && rm -rf /var/cache/apk/*
WORKDIR /usr/src/app
COPY . ./
RUN npm install
RUN npm run generate


FROM nginx:alpine

COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
