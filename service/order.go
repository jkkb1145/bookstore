package service

import (
	"demo02/repository"
)

type OrderService struct {
	orderDB *repository.OrderDAO
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderDB: repository.NewOrderDAO(),
	}
}

//func (o *OrderService) CreateOrder(req *controller.CreateOrderRequest)
