FROM golang:1.25-alpine

RUN apk add --no-cache git make

WORKDIR /app

ENV GOCACHE=/app/.cache/go-build
ENV GOMODCACHE=/app/.cache/go-mod

RUN go install github.com/air-verse/air@latest

EXPOSE 8080

CMD ["tail", "-f", "/dev/null"]