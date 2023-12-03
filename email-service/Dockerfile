FROM alpine as ca-certs

RUN apk add -U --no-cache ca-certificates

FROM golang:alpine as builder

WORKDIR /build

COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-cetificates.crt
COPY go.mod go.sum ./
COPY cmd/email/main.go ./

RUN go mod download

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o /app

FROM scratch as final

COPY --from=ca-certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/app /app

ENTRYPOINT ["/app"]