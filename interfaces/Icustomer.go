package interfaces

import 	"github.com/umamaheswari76/netxd_customer__dal/models"

type ICustomer interface {
	CreateCustomer(*models.Customer) (*models.CustomerResponse, error)
}
