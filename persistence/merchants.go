package persistence

import (
	"gopkg.in/mgo.v2"
)

type Merchant struct {
	// ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
}

type MerchantService interface {
	Register(merchant Merchant) (Merchant, error)
	All() ([]Merchant, error)
}

type MerchantServiceImp struct {
	Collection *mgo.Collection
}

func (m *MerchantServiceImp) Register(merchant Merchant) (Merchant, error) {
	err := m.Collection.Insert(&merchant)
	if err != nil {
		panic(err)
	}
	return merchant, nil
}

func (m *MerchantServiceImp) All() ([]Merchant, error) {
	var merchants []Merchant
	err := m.Collection.Find(nil).All(&merchants)
	if err != nil {
		panic(err)
	}
	return merchants, nil
}
