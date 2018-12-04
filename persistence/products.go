package persistence

import (
	"gopkg.in/mgo.v2"
)

type Product struct {
	ID     int64   `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Amount int16   `json:"amount"`
}

type ProductService interface {
}

type ProductServiceImp struct {
	Collection *mgo.Collection
}
