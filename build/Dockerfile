# voucher/vouch-proxy
# https://github.com/vouch/vouch-proxy
FROM golang:1.10 AS builder

WORKDIR ${GOPATH}/src/github.com/vouch/vouch-proxy

COPY . .

# RUN go-wrapper download  # "go get -d -v ./..."
# RUN ./do.sh build    # see `do.sh` for vouch build details
# RUN go-wrapper install # "go install -v ./..."

RUN ./do.sh goget && ./do.sh gobuildstatic && ./do.sh install

FROM builder AS dev

EXPOSE 9090

ENV TERM xterm

RUN apt-get update && apt-get install entr

CMD ["sh", "do.sh", "watch"]

FROM alpine:latest as prod

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY templates/ templates/
# see note for /static in main.go
COPY static /static
COPY do.sh /do.sh
COPY --from=builder /go/bin/vouch-proxy /vouch-proxy

EXPOSE 9090

ENTRYPOINT ["/vouch-proxy"]
HEALTHCHECK --interval=1m --timeout=5s CMD [ "/vouch-proxy", "-healthcheck" ]
