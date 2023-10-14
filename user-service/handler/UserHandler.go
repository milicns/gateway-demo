package handler

import (
	"context"
	"fmt"
	"user-service/model"
	pb "user-service/proto"
	"user-service/service"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	Service *service.UserService
}

func (handler *UserHandler) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("Hi from user")
	user := &model.User{
		Name:    request.Name,
		Surname: request.Surname,
	}

	err := handler.Service.Create(user)
	if err != nil {
		return &pb.CreateUserResponse{}, err
	}

	return &pb.CreateUserResponse{
		Message: "User created",
	}, nil
}