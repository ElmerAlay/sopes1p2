FROM golang:latest

WORKDIR /app

RUN go get -u github.com/gorilla/mux

COPY . .

RUN go build

CMD ["./app"]
