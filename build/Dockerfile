FROM golang:alpine AS base

WORKDIR /app/clusauth

COPY . .

FROM base AS dev

WORKDIR /app/clusauth

RUN go get -v -u github.com/cosmtrek/air

ENTRYPOINT [ "air" ]

FROM golang:alpine AS build

WORKDIR /app/clusauth

COPY --from=base /app/clusauth .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o clusauth .

FROM scratch AS prod

WORKDIR /app/clusauth

COPY configs .
COPY --from=build /app/clusauth/clusauth .

ENTRYPOINT [ "/app/clusauth/clusauth" ]
