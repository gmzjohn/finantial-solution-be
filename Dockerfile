FROM golang:1.25-alpine

RUN apk add --no-cache git make

WORKDIR /app

ENV GOPATH=/app/.go
ENV GOCACHE=/app/.go/cache

ENV PATH=$GOPATH/bin:$PATH

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]