FROM golang:1.17-alpine

WORKDIR /src
COPY ./src ./

RUN go build -o /echoserver

EXPOSE 8081

CMD ["/echoserver"]
