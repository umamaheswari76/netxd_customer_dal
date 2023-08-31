package models

type Transaction struct{
	Fromaccount   int `json:"fromaccount" bson:"fromaccount"`
	Toaccount     int `json:"toaccount" bson:"toaccount"`
	Amount        int `json:"amount" bson:"amount"`
}

// type TransactionResponse struct{

// }

// type Customer struct {
// 	CustomerId int32     `json:"customerid" bson:"customerid"`
// 	FirstName  string    `json:"firstname" bson:"firstname"`
// 	SecondName string    `json:"secondname" bson:"secondname"`
// 	BankId     string    `json:"bankid" bson:"bankid"`
// 	Balance    int32     `json:"balance" bson:"balance"`
// 	CreatedAt  time.Time `json:"createdat" bson:"createdat"`
// 	UpdatedAt  time.Time    `json:"updatedat" bson:"updatedat"`
// 	IsActive   string    `json:"isactive" bson:"isactive"`
// }