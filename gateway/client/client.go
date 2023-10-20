package client

import (
	"bytes"
	"context"
	"grpc-router/config"
	"net/http"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
)

type Client struct {
	refClient       *grpcreflect.Client
	cc              *grpc.ClientConn
	methods         map[string]config.MethodConfig
	descSource      grpcurl.DescriptorSource
	grpcServiceName string
}

func (client *Client) InvokeGrpcMethod(writer http.ResponseWriter, req *http.Request) {
	descSource := client.descSource
	mtdName := req.Context().Value("mtdName").(string)

	var resultBuffer bytes.Buffer

	rf, formatter, _ := grpcurl.RequestParserAndFormatterFor(grpcurl.Format("json"), descSource, false, true, req.Body)
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}
	err := grpcurl.InvokeRPC(context.Background(), descSource, client.cc, client.grpcServiceName+"/"+mtdName, nil, h, rf.Next)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultBuffer.Bytes())
}
