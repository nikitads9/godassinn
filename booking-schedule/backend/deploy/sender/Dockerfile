FROM golang:1.21-alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata
RUN apk add make

COPY ./booking-schedule/backend /github.com/nikitads9/godassinn/booking-schedule/backend/

WORKDIR /github.com/nikitads9/godassinn/booking-schedule/backend/

RUN make build-sender
RUN chown -R root ./bin/sender

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/bin .
COPY --from=builder /github.com/nikitads9/godassinn/booking-schedule/backend/configs/sender_config.yml .

CMD ["./sender", "-config", "sender_config.yml"]