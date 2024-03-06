FROM docker.io/library/golang:1.21-alpine3.17 AS build

## Build dependencies
RUN apk add --no-cache alpine-sdk

WORKDIR /build

## Build dependencies
RUN apk add --no-cache alpine-sdk

## Application dependencies
ADD go.mod go.sum ./
RUN go mod download


COPY cmd ./cmd
COPY internal ./internal

RUN cd cmd/app && go build

FROM docker.io/library/alpine:3.17 AS main

WORKDIR /app

COPY --from=build /build/cmd/app/app ./cmd/app/app

EXPOSE 80

CMD [ "./cmd/app/app" ]
