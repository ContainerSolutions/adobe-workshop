package main

import (
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type Endpoints struct {
	GetDealEndpoint endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		GetDealEndpoint: MakeGetDealEndpoint(s),
	}
}

func MakeGetDealEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getDealRequest)
		p, e := s.GetDeal(req.ID)
		return getDealResponse{Id: p.Id, Name: p.Name}, e
	}
}

type getDealRequest struct {
	ID int
}

type getDealResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
