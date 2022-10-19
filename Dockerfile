FROM golang:1.19 as base

WORKDIR /go/gotsrpc

COPY ./go.mod ./go.sum ./

RUN  go mod download -x

COPY . .

FROM base as builder

RUN GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o /service ./cmd/gotsrpc/gotsrpc.go

FROM alpine

COPY --from=builder /service /usr/bin/gotsrpc

CMD "/usr/bin/gotsrpc"

ENTRYPOINT "/usr/bin/gotsrpc"
