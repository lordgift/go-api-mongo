package persistence

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Merchant struct {
	// ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	BankAccount string `json:"bank_account" bson:"bankAccount"`
}

type MerchantService interface {
	Register(merchant Merchant) (Merchant, error)
	IsDuplicatedBankAccount(bankAccount string) (bool, error)
	All() ([]Merchant, error)
}

type MerchantServiceImp struct {
	Collection *mgo.Collection
}

func (m *MerchantServiceImp) IsDuplicatedBankAccount(bankAccount string) (bool, error) {
	result, err := m.Collection.Find(bson.M{"bankAccount": bankAccount}).Count()
	log.Printf("result:", result)
	return result > 0, err
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
