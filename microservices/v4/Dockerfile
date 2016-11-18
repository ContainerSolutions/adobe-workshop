FROM golang:1.6

COPY . /go/src/github.com/microservices-workshop/deals/

RUN go get github.com/gorilla/mux github.com/go-kit/kit/log github.com/go-kit/kit/endpoint github.com/go-kit/kit/transport/http gopkg.in/mgo.v2 github.com/go-kit/kit/metrics/prometheus github.com/prometheus/client_golang/prometheus

RUN go install github.com/microservices-workshop/deals

ENTRYPOINT /go/bin/deals

EXPOSE 8080