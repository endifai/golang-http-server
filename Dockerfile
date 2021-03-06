FROM golang:latest

RUN mkdir /app

ADD . /app 

WORKDIR /app

RUN go build -o server .

EXPOSE 5000

CMD ["/app/server"]
