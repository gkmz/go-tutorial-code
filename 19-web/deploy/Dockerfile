FROM golang:1.25.1-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /out/server ./graceful

FROM alpine:3.22
RUN adduser -D -u 10001 app
USER app
COPY --from=builder /out/server /server
EXPOSE 8082
ENTRYPOINT ["/server"]
