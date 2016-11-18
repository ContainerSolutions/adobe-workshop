package main

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/metrics"
	"time"
)

type Middleware func(Service) Service

func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) GetDeal(id int) (deal Deal, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "GetDeal",
			"id", id,
			"deal", deal.Id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.GetDeal(id)
}

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentingService(requestCount metrics.Counter, requestLatency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   requestCount,
		requestLatency: requestLatency,
		Service:        s,
	}
}

func (s *instrumentingService) GetDeal(id int) (deal Deal, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "GetDeal").Add(1)
		s.requestLatency.With("method", "GetDeal").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return s.Service.GetDeal(id)
}
