package server

import (
	"google.golang.org/grpc"
	db "kubeIT/database"
	"kubeIT/pkg/grpc/schema"
	"kubeIT/pkg/grpc/storage"
	"kubeIT/pkg/grpc/task"
	"kubeIT/pkg/grpc/user"
	"kubeIT/pkg/grpc/workflow"
	kubeitgrpc "kubeIT/server/grpc"
	"kubeIT/server/helpers"
	"kubeIT/server/s3handler"
	"log"
	"net"
)

type Api struct {
}

func (api *Api) Init(db *db.Database, authorizer *helpers.Authorizer, s3handler *s3handler.Api) {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	authserver := kubeitgrpc.NewUserManagerServer(db, authorizer)
	taskserver := kubeitgrpc.NewTaskManagementService(db, authorizer)
	schemaserver := kubeitgrpc.NewSchemaManagementService(db, authorizer)
	workflowserver := kubeitgrpc.NewWorkflowManagementService(db, authorizer)
	storageserver := kubeitgrpc.NewStorageService(db, authorizer, s3handler)

	// Register managers
	user.RegisterUserManagerServer(s, authserver)
	task.RegisterTaskManagementServiceServer(s, taskserver)
	schema.RegisterSchemaManagementServiceServer(s, schemaserver)
	workflow.RegisterWorkflowManagementServiceServer(s, workflowserver)
	storage.RegisterStorageManagementServiceServer(s, storageserver)

	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()
}
