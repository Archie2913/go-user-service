FROM golang:1.21

WORKDIR /app

COPY . .
COPY config/config.json /app/config/

ENV CONFIG_PATH=/app/config/config.json
ENV SERVER_PORT=:8080
ENV DATABASE_URL=postgres://prod-db:5432/myapp

RUN go build -o main ./cmd/api

CMD ["./main"] 