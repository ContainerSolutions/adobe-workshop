package main

import (
	"fmt"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"golang.org/x/net/context"
	"gopkg.in/mgo.v2"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	errc := make(chan error)
	ctx := context.Background()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC)
		logger = log.NewContext(logger).With("caller", log.DefaultCaller)
	}

	db, err := mgo.DialWithTimeout("deals-db:27017", time.Second*30)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	initData(db)

	fieldKeys := []string{"method"}
	var service Service
	{
		service = NewDealService(db, logger)
		service = LoggingMiddleware(logger)(service)
		service = NewInstrumentingService(
			kitprometheus.NewCounterFrom(
				stdprometheus.CounterOpts{
					Name:      "request_count",
					Help:      "Number of requests received.",
				},
				fieldKeys),
			kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
				Name:      "request_latency_microseconds",
				Help:      "Total duration of requests in microseconds.",
			}, fieldKeys),
			service,
		)
	}

	e := MakeEndpoints(service)

	// Create and launch the HTTP server.
	go func() {
		logger.Log("transport", "HTTP", "port", "8080")
		handler := MakeHTTPHandler(ctx, e, logger)
		errc <- http.ListenAndServe(":8080", handler)
	}()

	// Capture interrupts.
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("exit", <-errc)
}

// Initialize dummy data
func initData(s *mgo.Session) {
	c := s.DB("test").C("deals")
	err := c.Insert(&Deal{Id: 1, Name: "Buy 400 pairs, get one unmatched sock free!"},
		&Deal{Id: 2, Name: "Free shipping anywhere in the Andromeda Galaxy"})
	if err != nil {
		fmt.Printf("Error inserting records in database: %s\n", err.Error())
		panic(err)
	}
}
