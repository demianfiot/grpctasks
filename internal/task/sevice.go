package task

import (
	"context"
	"log"
	pb "rpcprac/pb/task/proto"
)

type TaskService struct {
	pb.UnimplementedTaskServiceServer
	Storage *TaskStorage
}

func (s *TaskService) CreateTask(
	ctx context.Context,
	req *pb.CreateTaskRequest,
) (*pb.CreateTaskResponse, error) {

	log.Println("CreateTask called with name:", req.Title)
	task := s.Storage.CreateTask(req.Title, req.Difficulty, req.XpReward)

	return &pb.CreateTaskResponse{
		Task: toPBTask(task),
	}, nil
}
func (s *TaskService) GetTask(
	ctx context.Context,
	req *pb.GetTaskRequest,
) (*pb.GetTaskResponse, error) {

	log.Println("GetTask called with id:", req.Id)
	task := s.Storage.GetTask(req.Id)
	return &pb.GetTaskResponse{
		Task: toPBTask(task),
	}, nil
}
func (s *TaskService) ListTasks(
	ctx context.Context,
	req *pb.ListTasksRequest,
) (*pb.ListTasksResponse, error) {

	tasks := s.Storage.ListTasks()

	pbUsers := make([]*pb.Task, 0, len(tasks))
	for _, task := range tasks {
		pbUsers = append(pbUsers, &pb.Task{
			Id:             task.Id,
			Title:          task.Title,
			Difficulty:     task.Difficulty,
			XpReward:       task.XpReward,
			AssignedUserId: task.AssignedUserId,
			Completed:      task.Completed,
		})
	}

	log.Printf("returning %d users", len(pbUsers))

	return &pb.ListTasksResponse{
		Tasks: pbUsers,
	}, nil
}
func (s *TaskService) AssignTask(
	ctx context.Context,
	req *pb.AssignTaskRequest,
) (*pb.AssignTaskResponse, error) {
	task := s.Storage.AssignTask(req.UserId, req.TaskId)
	return &pb.AssignTaskResponse{
		Task: toPBTask(task),
	}, nil
}
func (s *TaskService) CompleteTask(
	ctx context.Context,
	req *pb.CompleteTaskRequest,
) (*pb.CompleteTaskResponse, error) {
	task := s.Storage.CompleteTask(req.TaskId)
	return &pb.CompleteTaskResponse{
		Task: toPBTask(task),
	}, nil
}

func toPBTask(task *Task) *pb.Task {
	return &pb.Task{
		Id:             task.Id,
		Title:          task.Title,
		Difficulty:     task.Difficulty,
		XpReward:       task.XpReward,
		AssignedUserId: task.AssignedUserId,
		Completed:      task.Completed,
	}
}
