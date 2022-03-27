FROM golang:1.18-alpine as builder

WORKDIR /build

COPY go.mod .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main ./cmd/main.go

FROM alpine:latest

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait

RUN chmod +x /wait

COPY --from=builder main /bin/main

ENTRYPOINT ["/bin/sh", "-c", "/wait && /bin/main"]
