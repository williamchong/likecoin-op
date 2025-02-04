FROM --platform=$BUILDPLATFORM node:21.7.3-slim as builder

WORKDIR /usr/src/app
COPY . ./
RUN npm install
RUN npm run generate


FROM nginx:alpine

COPY --from=builder /usr/src/app/dist /usr/share/nginx/html
