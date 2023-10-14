package main

import (
	"context"
	"time"

	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(address string) *grpc.ClientConn {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cc *grpc.ClientConn
	var err error
	defer cancel()
	if cc, err = grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()); err != nil {
		panic(err)
	}
	return cc
}

func NewOrderClient(address string) (*grpcreflect.Client, *grpc.ClientConn) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var ccReflect *grpc.ClientConn
	var err error
	defer cancel()
	if ccReflect, err = grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()); err != nil {
		panic(err)
	}
	refClient := grpcreflect.NewClientAuto(context.Background(), ccReflect)
	return refClient, ccReflect
}
