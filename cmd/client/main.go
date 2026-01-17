package main

// import (
// 	"context"
// 	"log"
// 	taskpb "rpcprac/pb/task/proto"
// 	userpb "rpcprac/pb/user/proto"

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"
// )

// func main() {
// 	conn, err := grpc.NewClient(
// 		"localhost:50051",
// 		grpc.WithTransportCredentials(insecure.NewCredentials()),
// 	)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer conn.Close()

// 	userClient := userpb.NewUserServiceClient(conn)
// 	taskClient := taskpb.NewTaskServiceClient(conn)

// 	user, _ := userClient.CreateUser(
// 		context.Background(),
// 		&userpb.CreateUserRequest{Name: "Demian"},
// 	)

// 	task, _ := taskClient.CreateTask(
// 		context.Background(),
// 		&taskpb.CreateTaskRequest{
// 			Title:      "Clean room",
// 			Difficulty: "easy",
// 			XpReward:   10,
// 		},
// 	)

// 	taskClient.AssignTask(
// 		context.Background(),
// 		&taskpb.AssignTaskRequest{
// 			TaskId: task.Task.Id,
// 			UserId: user.User.Id,
// 		},
// 	)

// 	log.Println("DONE")
// }
