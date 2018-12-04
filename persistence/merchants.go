package persistence

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Merchant struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	Username    string        `json:"username"`
	Password    string        `json:"password"`
	Name        string        `json:"name"`
	BankAccount string        `json:"bank_account" bson:"bankAccount"`
	Products    []Product
}

type Product struct {
	Id     bson.ObjectId `bson:"_id,omitempty"`
	Name   string        `json:"name"`
	Price  float64       `json:"price"`
	Amount int16         `json:"amount"`
}

type MerchantService interface {
	Register(merchant Merchant) (Merchant, error)
	IsDuplicatedBankAccount(bankAccount string) (bool, error)
	FindById(id string) (Merchant, error)
	UpdateById(id string, merchant Merchant) (Merchant, error)

	All() ([]Merchant, error)
}

type MerchantServiceImp struct {
	Collection *mgo.Collection
}

func (m *MerchantServiceImp) IsDuplicatedBankAccount(bankAccount string) (bool, error) {
	count, err := m.Collection.Find(bson.M{"bankAccount": bankAccount}).Count()
	return count > 0, err
}

func (m *MerchantServiceImp) Register(merchant Merchant) (Merchant, error) {
	err := m.Collection.Insert(&merchant)
	if err != nil {
		panic(err)
	}
	return merchant, nil
}

func (m *MerchantServiceImp) FindById(id string) (Merchant, error) {
	var merchant Merchant
	err := m.Collection.FindId(bson.ObjectIdHex(id)).One(&merchant)
	if err != nil {
		panic(err)
	}
	return merchant, nil
}

func (m *MerchantServiceImp) UpdateById(id string, merchant Merchant) (Merchant, error) {
	// err := m.Collection.FindId(bson.ObjectIdHex(id)).One(&merchant)
	err := m.Collection.UpdateId(bson.ObjectIdHex(id), merchant)
	if err != nil {
		panic(err)
	}
	return merchant, err
}

func (m *MerchantServiceImp) All() ([]Merchant, error) {
	var merchants []Merchant
	err := m.Collection.Find(nil).All(&merchants)
	if err != nil {
		panic(err)
	}
	return merchants, nil
}
