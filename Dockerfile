FROM golang:1.24.5-alpine AS builder

RUN apk update && apk add ca-certificates git gcc g++ libc-dev binutils

WORKDIR /opt

COPY . .

RUN go mod download && go mod verify

RUN go build -o bin/app ./cmd/task_manager

FROM alpine:3.19 AS runner

RUN apk update && apk add ca-certificates libc6-compat openssh bash && rm -rf /var/cache/apk/*

WORKDIR /opt

COPY .env /opt
COPY --from=builder /opt/bin/app ./

CMD ["./app"]