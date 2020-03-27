# voucher/vouch-proxy
# https://github.com/vouch/vouch-proxy
FROM golang:1.13.6-alpine3.11 AS builder

WORKDIR /go/src/github.com/vouch/vouch-proxy

COPY . /go/src/github.com/vouch/vouch-proxy/

# RUN go-wrapper download  # "go get -d -v ./..."
# RUN ./do.sh build    # see `do.sh` for vouch build details
# RUN go-wrapper install # "go install -v ./..."

RUN apk update && apk upgrade && apk add --no-cache bash git openssh make && \
  make build install

FROM alpine:3.11.3

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY templates/ templates/
# see note for /static in main.go
COPY static /static
COPY do.sh /do.sh
COPY --from=builder /go/bin/vouch-proxy /vouch-proxy

EXPOSE 9090

ENTRYPOINT ["/vouch-proxy"]
HEALTHCHECK --interval=1m --timeout=5s CMD [ "/vouch-proxy", "-healthcheck" ]
