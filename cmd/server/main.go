package main

import (
	"log"
	"net"
	"rpcprac/handler"
	taskpb "rpcprac/pb/task/proto"
	userpb "rpcprac/pb/user/proto"
	"rpcprac/repository"
	"rpcprac/service"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	dbConfig := repository.DBConfigFromViper()
	bd, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository := repository.NewRepository(bd)
	services := service.NewService(repository)
	userHandler := handler.NewUserHandler(services.User)
	taskHandler := handler.NewTaskHandler(services.Task)

	userpb.RegisterUserServiceServer(grpcServer, userHandler)
	taskpb.RegisterTaskServiceServer(grpcServer, taskHandler)

	log.Println("gRPC server started on :50051")
	reflection.Register(grpcServer)

	log.Fatal(grpcServer.Serve(lis))
}
