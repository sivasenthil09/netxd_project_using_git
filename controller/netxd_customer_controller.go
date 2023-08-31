package controllers

import (
	"context"

	"github.com/sivasenthil09/netxd_dal/netxd_customer_dal/interfaces"
	"github.com/sivasenthil09/netxd_dal/netxd_customer_dal/models"
	pro "github.com/sivasenthil09/netxd_project_proto"
)

type RPCserver struct {
	pro.UnimplementedCustomerServiceServer
}

var (
	CustomerService interfaces.ICustomer
)

func (s *RPCserver) CreateCustomer(ctx context.Context, req *pro.Customer) (*pro.CustomerResponse, error) {
	dbCustomer := &models.Customer{
		CustomerID: req.CustomerID}
	result, err := CustomerService.CreateCustomer(dbCustomer)
	if err != nil {
		return nil, err
	} else {
		responseCustomer := &pro.CustomerResponse{
			CustomerID: result.CustomerID,
		}
		return responseCustomer, nil
	}
}
