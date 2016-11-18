package main

import (
	"github.com/go-kit/kit/log"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Service interface {
	GetDeal(id int) (Deal, error)
}

type Deal struct {
	Id   int    `json:"id" bson:"id"`
	Name string `json:"name" bson:"name"`
}

func NewDealService(db *mgo.Session, logger log.Logger) Service {
	return &dealService{
		db:     *db,
		logger: logger,
	}
}

type dealService struct {
	db     mgo.Session
	logger log.Logger
}

func (s *dealService) GetDeal(id int) (Deal, error) {
	c := s.db.DB("test").C("deals")
	r := Deal{}
	err := c.Find(bson.M{"id": id}).One(&r)
	if err != nil {
		return r, err
	}
	return r, nil
}
