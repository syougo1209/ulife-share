FROM golang:1.17.1-alpine as builder

WORKDIR /myapp
COPY ./backend/go.mod ./backend/go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@v1.27.3
COPY ./backend  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' -o /build/main ./cmd/ulife-share
ENTRYPOINT ["air", "-c", ".air.toml"]
