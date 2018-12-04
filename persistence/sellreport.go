package persistence

import (
	"gopkg.in/mgo.v2"
)

type Sell struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int16   `json:"amount"`
}

type SellService interface {
}

type SellServiceImp struct {
	Collection *mgo.Collection
}
