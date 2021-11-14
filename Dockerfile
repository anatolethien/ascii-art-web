FROM golang:1.16-bullseye

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o bin/ascii-art-web main.go

CMD ["./bin/ascii-art-web"]
