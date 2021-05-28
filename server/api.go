package server

import (
	"google.golang.org/grpc"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/user"
	kubeitgrpc "kubeIT/server/grpc"
	"log"
	"net"
)

type Api struct {
}

func (api *Api) Init(db *db.Database) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	authserver := kubeitgrpc.NewUserManagerServer(db)
	// Register usermanager
	user.RegisterUserManagerServer(s, authserver)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
}
