FROM node:18.16 AS web-builder
WORKDIR /app
COPY . .
RUN npm ci --prefix web
RUN make build_web

FROM golang:1.20-alpine AS core-builder
RUN apk add --no-cache make gcc musl-dev
WORKDIR /app
COPY --from=web-builder /app .
RUN make build_core

FROM alpine:latest
WORKDIR /app
COPY --from=core-builder /app/qlookout ./
ENTRYPOINT ["./qlookout"]

