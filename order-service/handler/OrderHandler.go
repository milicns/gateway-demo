package handler

import (
	"context"
	"fmt"
	"order-service/model"
	pb "order-service/proto"
	"order-service/service"
)

type OrderHandler struct {
	Service *service.OrderService
	pb.UnimplementedOrderServiceServer
}

func (handler *OrderHandler) CreateOrder(ctx context.Context, request *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	fmt.Println("Hi from order")
	order := &model.Order{
		User: request.User,
	}

	err := handler.Service.Create(order)
	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}

	return &pb.CreateOrderResponse{
		Message: "Order created",
	}, nil
}

func (handler *OrderHandler) CreateOrderV2(ctx context.Context, request *pb.CreateOrderV2Request) (*pb.CreateOrderResponse, error) {
	return &pb.CreateOrderResponse{
		Message: "Create order v2",
	}, nil
}
