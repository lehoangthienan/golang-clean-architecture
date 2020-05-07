FROM golang:1.12-alpine as builder

RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates

ENV GO111MODULE=on \
    GOPROXY=https://proxy.golang.org

WORKDIR /github.com/lehoangthienan/marvel-heroes-backend

COPY . .

RUN go mod tidy
RUN go mod download
RUN go mod verify

RUN go build -o marvel-heroes-backend /github.com/lehoangthienan/marvel-heroes-backend/cmd/server/main.go

EXPOSE 4002

CMD ["./marvel-heroes-backend"]