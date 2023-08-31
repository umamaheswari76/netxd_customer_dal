package services

import (
	"context"
	"github.com/umamaheswari76/netxd_customer_/interfaces"
	"github.com/umamaheswari76/netxd_customer__dal/models"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}

func InitializeCustomerService(collection *mongo.Collection, ctx context.Context) interfaces.ICustomer {
	return &CustomerService{collection, ctx}
}

// type Customer struct {
//     CustomerId int32     `json:"customerid" bson:"customerid"`
//     FirstName  string    `json:"firstname" bson:"firstname"`
//     SecondName string    `json:"secondname" bson:"secondname"`
//     BankId     string    `json:"bankid" bson:"bankid"`
//     Balance    int32     `json:"balance" bson:"balance"`
//     CreatedAt  time.Time `json:"createdat" bson:"createdat"`
//     UpdatedAt  time.Time `json:"updatedat" bson:"updatedat"`
//     IsActive   string    `json:"isactive" bson:"isactive"`
// }

func (c *CustomerService) CreateCustomer(customer *models.Customer) (*models.CustomerResponse, error){
	customer.CreatedAt = time.Now()
	customer.UpdatedAt = customer.CreatedAt
	customer.IsActive = "yes"

	_, err := c.CustomerCollection.InsertOne(c.ctx, &customer)
	if err != nil{
		return nil, err
	}
	
	return &models.CustomerResponse{
		CustomerId: customer.CustomerId,
		CreatedAt:  customer.CreatedAt,
	},nil
}
