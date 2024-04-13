FROM golang:1.21-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

COPY ./booking-schedule/backend /github.com/nikitads9/godassinn/booking-schedule/backend/

WORKDIR /github.com/nikitads9/godassinn/booking-schedule/backend/

RUN make build-bookings
RUN chown -R root ./bin/bookings

FROM alpine:latest

WORKDIR /root/

RUN apk add nginx
ADD ./deploy/nginx/nginx.conf /etc/nginx/nginx.conf

COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/bin .
COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/configs/booking_config.yml .

CMD [ "./bookings", "-config", "booking_config.yml", "-keyfile", "/etc/ssl/certs/key.pem", "-certfile", "/etc/ssl/certs/cert.pem" ]