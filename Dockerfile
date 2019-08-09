FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go build -o hihello .

CMD ["/app/hihello"]

EXPOSE 8082
