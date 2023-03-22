### Builder Image
FROM golang:1.19-alpine as builder
WORKDIR /workspace

RUN apk update
RUN apk add --no-cache ca-certificates openssl bash make

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

### Release Image
FROM alpine:3.17
WORKDIR /workspace

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /workspace/bin/* ./

ENTRYPOINT ["./api"]
