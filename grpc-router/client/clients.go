package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	RefClient *grpcreflect.Client
	Cc        *grpc.ClientConn
}

type ClientRegistry struct {
	Clients map[string]Client
}

func (registry *ClientRegistry) NewClient(address string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var cc *grpc.ClientConn
	var err error
	defer cancel()
	if cc, err = grpc.DialContext(ctx, address,
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock()); err != nil {
		panic(err)
	}
	refClient := grpcreflect.NewClientAuto(context.Background(), cc)
	registry.Clients[address] = Client{RefClient: refClient, Cc: cc}
}

func (client *Client) ListServices(writer http.ResponseWriter, req *http.Request) {
	fmt.Println("hi")

	sourceReflect := grpcurl.DescriptorSourceFromServer(context.Background(), client.RefClient)
	list, err := grpcurl.ListServices(sourceReflect)
	dsc, err := sourceReflect.FindSymbol("UserService")
	sd, _ := dsc.(*desc.ServiceDescriptor)
	mtd := sd.FindMethodByName("CreateUser")
	fmt.Println(mtd.GetInputType())
	if err != nil {
		return
	}
	fmt.Println(list)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
}

func (client *Client) InvokeGrpcMethod(writer http.ResponseWriter, req *http.Request) {

	sourceReflect := grpcurl.DescriptorSourceFromServer(context.Background(), client.RefClient)
	var resultBuffer bytes.Buffer

	rf, formatter, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), sourceReflect, false, true, req.Body)
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}

	err := grpcurl.InvokeRPC(context.Background(), sourceReflect, client.Cc, "UserService/CreateUser", nil, h, rf.Next)
	if err != nil {
		fmt.Print(err)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultBuffer.Bytes())
}

func parseSymbol(svcAndMethod string) (string, string) {
	pos := strings.LastIndex(svcAndMethod, "/")
	if pos < 0 {
		pos = strings.LastIndex(svcAndMethod, ".")
		if pos < 0 {
			return "", ""
		}
	}
	return svcAndMethod[:pos], svcAndMethod[pos+1:]
}
