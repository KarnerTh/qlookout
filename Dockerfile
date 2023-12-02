FROM golang:1.21-alpine AS builder
RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN make build_project

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/qlookout ./
ENTRYPOINT ["./qlookout"]

