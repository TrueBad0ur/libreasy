FROM golang:1.21.6-alpine3.19 as builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o /build/libreasy main.go

FROM alpine:3.19
WORKDIR /app

COPY --from=builder /build/libreasy libreasy
COPY assets assets
COPY css css
COPY templates templates

ENTRYPOINT ["/app/libreasy"]
EXPOSE 8080