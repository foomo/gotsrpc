FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -u 1001 -g 1001 gotsrpc

COPY gotsrpc /usr/bin/

USER gotsrpc
WORKDIR /home/gotsrpc

ENTRYPOINT ["gotsrpc"]
