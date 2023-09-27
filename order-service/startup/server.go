package startup

import (
	"fmt"
	"log"
	"net"
	"order-service/handler"
	pb "order-service/proto"
	"order-service/repo"
	"order-service/service"
	"order-service/startup/config"
	"os"
	"os/signal"
	"syscall"

	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server *Server) Start() {
	mongoClient := server.initMongoClient()
	orderRepo := server.initOrderRepo(mongoClient)
	orderService := server.initOrderService(orderRepo)
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &handler.OrderHandler{Service: orderService})
	reflection.Register(grpcServer)
	if lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("ORDER_SERVICE_PORT"))); err != nil {
		panic(err)
	} else {
		go grpcServer.Serve(lis)
	}
	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh
	grpcServer.Stop()
}

func (server *Server) initMongoClient() *mongo.Client {
	client, err := repo.GetClient(server.config.OrdersDBHost, server.config.OrdersDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initOrderRepo(client *mongo.Client) *repo.OrderRepository {
	store := repo.NewOrderRepository(client)
	return store
}

func (server *Server) initOrderService(repo *repo.OrderRepository) *service.OrderService {
	service := service.NewOrderService(repo)
	return service
}
