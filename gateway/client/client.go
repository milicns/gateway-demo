package client

import (
	"bytes"
	"context"
	"fmt"
	"gateway/config"
	"io"
	"net/http"

	"github.com/fullstorydev/grpcurl"
	"github.com/goccy/go-json"
	"github.com/gorilla/mux"
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
	reader := prepareReader(req)
	fmt.Println(req.Header)
	var resultBuffer bytes.Buffer
	headers := make([]string, 0, len(req.Header))
	req.Header["Content-Type"] = []string{"application/grpc"}
	for k, v := range req.Header {
		headers = append(headers, fmt.Sprintf("%s: %s", k, v))
	}
	fmt.Println(headers)
	rf, formatter, _ := grpcurl.RequestParserAndFormatter(grpcurl.Format("json"), descSource, reader, grpcurl.FormatOptions{})
	h := &grpcurl.DefaultEventHandler{
		Out:            &resultBuffer,
		Formatter:      formatter,
		VerbosityLevel: 0,
	}
	err := grpcurl.InvokeRPC(context.Background(), descSource, client.cc, client.grpcServiceName+"/"+mtdName, headers, h, rf.Next)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	writer.Write(resultBuffer.Bytes())
}

func prepareReader(req *http.Request) io.Reader {
	var reader io.Reader
	params := mux.Vars(req)
	var reqBodyMap map[string]interface{}
	json.NewDecoder(req.Body).Decode(&reqBodyMap)
	if reqBodyMap != nil {
		for k, v := range params {
			reqBodyMap[k] = v
		}
		reqBodyBytes, _ := json.Marshal(reqBodyMap)
		reader = bytes.NewReader(reqBodyBytes)
	} else {
		paramsBytes, _ := json.Marshal(params)
		reader = bytes.NewReader(paramsBytes)
	}

	return reader

}
