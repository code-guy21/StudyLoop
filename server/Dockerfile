FROM golang:1.21-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:3.18

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/configs ./configs

RUN adduser -D tutorlink

USER tutorlink

EXPOSE 8080

CMD ["./main"]


