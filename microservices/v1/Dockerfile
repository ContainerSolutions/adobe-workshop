FROM golang:1.6

COPY . /go/src/github.com/microservices-workshop/deals

RUN go install github.com/microservices-workshop/deals

ENTRYPOINT /go/bin/deals

EXPOSE 80