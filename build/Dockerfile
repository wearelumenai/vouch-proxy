# voucher/vouch-proxy
# https://github.com/vouch/vouch-proxy
FROM golang:1.13.6-alpine3.11 AS base

WORKDIR /go/src/github.com/vouch/vouch-proxy

COPY . .

# RUN go-wrapper download  # "go get -d -v ./..."
# RUN ./do.sh build    # see `do.sh` for vouch build details
# RUN go-wrapper install # "go install -v ./..."

RUN apk update && apk upgrade && apk add --no-cache bash git openssh make && \
  make build install

# dev
FROM base AS dev

RUN go get -v -u github.com/cosmtrek/air

EXPOSE 8003

ENTRYPOINT [ "air" ]

# build
FROM base AS build

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vouch-proxy .

# prod
FROM alpine:3.11.3 AS prod

COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY templates/ templates/
# see note for /static in main.go
COPY static /static
COPY do.sh /do.sh
COPY --from=build /go/src/github.com/vouch/vouch-proxy/vouch-proxy /vouch-proxy

EXPOSE 8003

ENTRYPOINT ["/vouch-proxy"]

HEALTHCHECK --interval=1m --timeout=5s CMD [ "/vouch-proxy", "-healthcheck" ]
