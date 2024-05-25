FROM golang:1.22-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext

# зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY cmd ./
COPY configs ./
COPY internal ./
COPY schema ./

RUN go build -o ./bin/runner main.go