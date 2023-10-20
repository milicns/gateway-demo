package startup

import (
	"context"
	"fmt"
	"grpc-router/client"
	"grpc-router/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

func (s *Server) Start() {
	clientRegistry := s.prepareClients()
	router := s.prepareRoutes(clientRegistry)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.config.Gateway.Port), router))
}

func (s *Server) prepareClients() *client.ClientRegistry {
	clientRegistry := &client.ClientRegistry{
		Clients: make(map[string]client.Client),
	}

	for _, s := range s.config.Services {
		clientRegistry.NewClient(s.Address, s.Methods)
	}

	return clientRegistry
}
func (s *Server) prepareRoutes(clientRegistry *client.ClientRegistry) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, srv := range s.config.Services {
		client := clientRegistry.Clients[srv.Address]
		for mtdName, mtdConf := range srv.Methods {
			path := s.config.Gateway.Route + srv.ServiceRoute + mtdConf.MethodRoute
			router.HandleFunc("/"+path, methodNameMiddleware(mtdName, client.InvokeGrpcMethod)).Methods(mtdConf.Type)
		}
	}
	return router
}

func methodNameMiddleware(mtdName string, h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "mtdName", mtdName)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
