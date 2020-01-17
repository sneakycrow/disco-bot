FROM golang:1.13.6

RUN mkdir /app
ADD . /app

WORKDIR /app

RUN go build

CMD ["/app/disco-bot"]

