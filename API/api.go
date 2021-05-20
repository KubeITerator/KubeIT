package API

import (
	kubeitmodel "KubeIT-gRPC/model/grpc"
	"KubeIT-gRPC/API/gateway"
	kubeitservice "KubeIT-gRPC/API/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

type Api struct {
}

func (api *Api) Init{
	ctx := context.Background()
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	authserver := kubeitservice.NewUserManagerServer()
	// Attach the Greeter API to the server
	kubeitmodel.RegisterUserManagerServer(s, authserver)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		ctx,
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	gateway.HandleAuth(ctx, gwmux)

	err = kubeitmodel.RegisterUserManagerHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8091",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8091")
	log.Fatalln(gwServer.ListenAndServe())
}
