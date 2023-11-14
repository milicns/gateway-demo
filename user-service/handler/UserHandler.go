package handler

import (
	"context"
	"fmt"
	"user-service/model"
	pb "user-service/proto"
	"user-service/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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
		return &pb.CreateUserResponse{}, status.Error(codes.Internal, "Already exists")
	}

	return &pb.CreateUserResponse{
		Message: "User created",
	}, nil
}

func (handler *UserHandler) PathParamsTest(ctx context.Context, request *pb.PathParamsRequest) (*pb.PathParamResponse, error) {
	return &pb.PathParamResponse{
		Message: "Param1: " + request.Id + " Param2: " + request.Name,
	}, nil
}

func (handler *UserHandler) PathParamTest(ctx context.Context, request *pb.PathParamRequest) (*pb.PathParamResponse, error) {
	return &pb.PathParamResponse{
		Message: "Param1: " + request.Id,
	}, nil
}

func (handler *UserHandler) PathParamsAndBodyTest(ctx context.Context, request *pb.PathParamsAndBodyRequest) (*pb.PathParamResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	return &pb.PathParamResponse{
		Message: "Param1: " + request.Id + " Param2: " + request.Name + ". Body: " + request.Surname + "," + fmt.Sprint(request.Number),
	}, nil
}
