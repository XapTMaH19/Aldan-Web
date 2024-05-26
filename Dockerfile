FROM golang:1.22.2-alpine AS builder

WORKDIR /usr/local/src

RUN apk --no-cache add bash git make gcc gettext musl-dev

# зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

# build
COPY cmd ./cmd
COPY configs ./configs
COPY internal ./internal
COPY schema ./scheme
COPY server.go ./

RUN go build -o ./bin/runner cmd/main.go

FROM alpine AS runner

COPY --from=builder /usr/local/src/bin/runner /
COPY configs ./configs
COPY .env ./

CMD ["/runner"]