FROM golang:1.19.2-alpine3.16 as builder

WORKDIR /app

COPY apps/gateway/go.sum apps/gateway/go.mod ./

RUN go mod download

COPY apps/gateway ./

RUN CGO_ENABLED=0 go build -o /gateway

FROM alpine:latest

WORKDIR /app

COPY --from=builder /gateway ./gateway

CMD [ "./gateway" ]
