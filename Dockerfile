FROM golang:1.16-alpine

WORKDIR /ascii-art-web

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/ascii-art-web cmd/main.go

CMD ["./bin/ascii-art-web"]
