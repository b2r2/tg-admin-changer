# syntax=docker/dockerfile:1.3
FROM golang:1.17.1-alpine3.14 as build

WORKDIR /usr/local/app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -ldflags "-s -w" -o tg-admin-changer ./cmd/tg-admin-changer/main.go


FROM alpine:3.14 as tg-admin-changer

RUN sed -i 's/https\:\/\/dl-cdn.alpinelinux.org/http\:\/\/mirror.clarkson.edu/g' /etc/apk/repositories && apk add ca-certificates --no-cache

WORKDIR /usr/local/app

RUN --mount=type=secret,id=TOKEN \
    export TOKEN=$(cat /run/secrets/TOKEN) && \
    echo $TOKEN >> .env

COPY --from=build /usr/local/app/tg-admin-changer /bin/tg-admin-changer

CMD /bin/tg-admin-changer