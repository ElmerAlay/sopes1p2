FROM golang:latest

WORKDIR /app

RUN go get -u github.com/gorilla/mux

COPY . .

CMD ["./app"]