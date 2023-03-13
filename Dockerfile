FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY *.go ./
COPY client/ ./

RUN go mod download

RUN go build -o /app/bin/server

CMD ["/app/bin/server" ]
