FROM golang:1.24-alpine

RUN apk add --no-cache git curl ca-certificates

RUN go install github.com/air-verse/air@latest

WORKDIR /app
COPY . .

RUN go mod download

EXPOSE 8080

CMD ["air"]
