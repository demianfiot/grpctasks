package main

import (
	"log"
	"net"
	"rpcprac/internal/task"
	"rpcprac/internal/user"
	taskpb "rpcprac/pb/task/proto"
	userpb "rpcprac/pb/user/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	userStorage := user.NewUserStorage()
	taskStorage := task.NewTaskStorage()

	userpb.RegisterUserServiceServer(
		grpcServer,
		&user.UserService{Storage: userStorage},
	)

	taskpb.RegisterTaskServiceServer(
		grpcServer,
		&task.TaskService{Storage: taskStorage},
	)

	log.Println("gRPC server started on :50051")
	reflection.Register(grpcServer)

	log.Fatal(grpcServer.Serve(lis))
}
