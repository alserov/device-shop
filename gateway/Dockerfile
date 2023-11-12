FROM alpine as root-certs

RUN apk add -U --no-cache ca-certificates

FROM golang:alpine as builder

WORKDIR /build

COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY go.mod go.sum ./
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./gateway

FROM scratch as final

COPY --from=builder /build/api /api
COPY --from=root-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/api"]