FROM golang:1.24-alpine

RUN apk add --no-cache git curl ca-certificates

RUN go install github.com/air-verse/air@latest

WORKDIR /app

# 🔁 copy only go.mod ก่อนเพื่อ cache ได้
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# 🔁 copy source ภายหลัง
COPY . .

EXPOSE 8080

CMD ["air"]
