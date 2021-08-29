# syntax=docker/dockerfile:1
FROM golang:1.17.0-alpine as builder

RUN apk add git
WORKDIR /app
COPY . ./
RUN ./gomake build -L debug && cp dist/checksum-linux-amd64/checksum /checksum


FROM alpine
COPY --from=builder /checksum /usr/bin/checksum
ENTRYPOINT [ "checksum" ]
