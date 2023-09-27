package service

import (
	"errors"
	"order-service/model"
	"order-service/repo"
)

type OrderService struct {
	Repo *repo.OrderRepository
}

func NewOrderService(repo *repo.OrderRepository) *OrderService {
	return &OrderService{
		Repo: repo,
	}
}

func (service *OrderService) Create(order *model.Order) error {
	_, err := service.GetOne(order.User)
	if err == nil {
		return errors.New("user with this name already created order")
	}
	return service.Repo.Create(order)
}

func (service *OrderService) GetOne(user string) (*model.Order, error) {
	return service.Repo.GetOne(user)
}
