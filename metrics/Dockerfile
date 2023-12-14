FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
COPY . .
COPY cmd/metrics/main.go ./
COPY config/local.yaml ./

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app

FROM scratch as final

COPY --from=builder /build/app /app
COPY --from=builder /build/config /config

ENTRYPOINT ["/app"]