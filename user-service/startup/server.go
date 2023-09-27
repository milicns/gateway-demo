package startup

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"user-service/handler"
	pb "user-service/proto"
	"user-service/repo"
	"user-service/service"
	"user-service/startup/config"

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
	userRepo := server.initUserStore(mongoClient)
	userService := server.initUserService(userRepo)
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &handler.UserHandler{Service: userService})
	reflection.Register(grpcServer)
	if lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("USER_SERVICE_PORT"))); err != nil {
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
	client, err := repo.GetClient(server.config.UsersDBHost, server.config.UsersDBPort)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func (server *Server) initUserStore(client *mongo.Client) *repo.UserRepository {
	store := repo.NewUserRepository(client)
	return store
}

func (server *Server) initUserService(repo *repo.UserRepository) *service.UserService {
	service := service.NewUserService(repo)
	return service
}
