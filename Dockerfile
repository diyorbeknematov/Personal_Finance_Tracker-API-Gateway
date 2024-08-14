FROM golang:1.22.2 AS builder

WORKDIR /api-gateway

COPY . ./
RUN go mod download

COPY config/model.conf ./
COPY .env .

RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd -a -installsuffix cgo -o ./../gateway .

FROM alpine:latest

WORKDIR /api-gateway

COPY --from=builder /api-gateway/gateway .
COPY --from=builder /api-gateway/pkg/logs/app.log ./pkg/logs/
COPY --from=builder /api-gateway/config/model.conf ./config/
COPY --from=builder /api-gateway/.env .

EXPOSE 9999

CMD [ "./gateway" ]