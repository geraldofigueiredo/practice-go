FROM golang:1.18-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY sender/main.go .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o sender .

FROM scratch

COPY --from=builder ["/build/sender", "/"]

ENTRYPOINT ["/sender"]