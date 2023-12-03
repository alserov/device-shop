FROM alpine as ca-certs

RUN apk add -U --no-cache ca-certificates

FROM golang:alpine as builder

WORKDIR /build

COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY go.mod go.sum ./
COPY *.go ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./service

FROM scratch as final

COPY --from=builder /build/service /service
COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["/service"]