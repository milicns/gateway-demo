package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
)

type OrderInvoker struct {
	refClient *grpcreflect.Client
	cc        *grpc.ClientConn
}

type UserInvoker struct {
	refClient *grpcreflect.Client
	cc        *grpc.ClientConn
}

func (userInvoker *UserInvoker) invokeCreateUser(writer http.ResponseWriter, req *http.Request) {

	sourceReflect := grpcurl.DescriptorSourceFromServer(context.Background(), userInvoker.refClient)
	var resultBuffer bytes.Buffer

	rf, formatter, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), sourceReflect, false, true, req.Body)
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}

	err := grpcurl.InvokeRPC(context.Background(), sourceReflect, userInvoker.cc, "UserService/CreateUser", nil, h, rf.Next)
	if err != nil {
		fmt.Print(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultBuffer.Bytes())
}

func (orderInvoker *OrderInvoker) invokeCreateOrder(writer http.ResponseWriter, req *http.Request) {

	sourceReflect := grpcurl.DescriptorSourceFromServer(context.Background(), orderInvoker.refClient)
	var resultBuffer bytes.Buffer

	rf, formatter, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), sourceReflect, false, true, req.Body)
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}

	err := grpcurl.InvokeRPC(context.Background(), sourceReflect, orderInvoker.cc, "OrderService/CreateOrder", nil, h, rf.Next)
	if err != nil {
		fmt.Print(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultBuffer.Bytes())
}
