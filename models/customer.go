package models

import "time"

type Customer struct {
	CustomerId int32     `json:"customerid" bson:"customerid"`
	FirstName  string    `json:"firstname" bson:"firstname"`
	SecondName string    `json:"secondname" bson:"secondname"`
	BankId     string    `json:"bankid" bson:"bankid"`
	Balance    int32     `json:"balance" bson:"balance"`
	CreatedAt  time.Time `json:"createdat" bson:"createdat"`
	UpdatedAt  time.Time    `json:"updatedat" bson:"updatedat"`
	IsActive   string    `json:"isactive" bson:"isactive"`
}

type CustomerResponse struct {
	CustomerId int32  `json:"customerid" bson:"customerid"`
	CreatedAt  time.Time `json:"createdat" bson:"createdat"`
}
