package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"grpc-router/config"
	"net/http"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	RefClient  *grpcreflect.Client
	Cc         *grpc.ClientConn
	Methods    map[string]config.MethodConfig
	DescSource grpcurl.DescriptorSource
}

type ClientRegistry struct {
	Clients map[string]Client
}

func (registry *ClientRegistry) NewClient(address string, methods map[string]config.MethodConfig) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cc *grpc.ClientConn
	var err error
	defer cancel()
	if cc, err = grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()); err != nil {
		panic(err)
	}
	refClient := grpcreflect.NewClientAuto(context.Background(), cc)
	descSource := grpcurl.DescriptorSourceFromServer(context.Background(), refClient)
	registry.Clients[address] = Client{RefClient: refClient, Cc: cc, Methods: methods, DescSource: descSource}
}

func (client *Client) ListServices(writer http.ResponseWriter, req *http.Request) {
	descSource := grpcurl.DescriptorSourceFromServer(context.Background(), client.RefClient)
	services, _ := descSource.ListServices()
	dsc, _ := descSource.FindSymbol(services[0])
	sd, _ := dsc.(*desc.ServiceDescriptor)
	fmt.Println(sd.GetMethods()[0].GetName())

	//rf, _, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), sourceReflect, false, true, req.Body)
	//dsc, err := sourceReflect.FindSymbol("UserService")
	//sd, _ := dsc.(*desc.ServiceDescriptor)
	//mtd := sd.FindMethodByName("CreateUser")

	//msgFactory := dynamic.NewMessageFactoryWithDefaults()
	//r := msgFactory.NewMessage(mtd.GetInputType())
	//rf.Next(r)
	//fmt.Println(r)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (client *Client) InvokeGrpcMethod(writer http.ResponseWriter, req *http.Request) {
	//param := mux.Vars(req)
	//fmt.Println(param)
	descSource := client.DescSource
	services, _ := descSource.ListServices()
	fmt.Println(req.Body)
	mtdName := req.Context().Value("mtdName").(string)

	var resultBuffer bytes.Buffer

	rf, formatter, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), descSource, false, true, req.Body)
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}
	err := grpcurl.InvokeRPC(context.Background(), descSource, client.Cc, services[0]+"/"+mtdName, nil, h, rf.Next)
	if err != nil {
		fmt.Print(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(resultBuffer.Bytes())
}
