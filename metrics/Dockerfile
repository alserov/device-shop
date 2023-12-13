FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -C ./cmd/metrics/

FROM scratch as final

COPY --from=builder /build/cmd/metrics /app

ENTRYPOINT ["/app"]