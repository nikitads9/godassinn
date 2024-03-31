FROM golang:1.21-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

COPY ./booking-schedule/backend /github.com/nikitads9/godassinn/booking-schedule/backend/

WORKDIR /github.com/nikitads9/godassinn/booking-schedule/backend/

RUN make build-auth
RUN chown -R root ./bin/auth

FROM alpine:latest

RUN apk add nginx
ADD ./booking-schedule/backend/deploy/nginx/nginx.conf /etc/nginx/nginx.conf

WORKDIR /root/

COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/bin .
COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/configs/auth_config.yml .

CMD [ "./auth", "-config", "auth_config.yml" , "-keyfile", "/etc/ssl/certs/key.pem", "-certfile", "/etc/ssl/certs/cert.pem" ]