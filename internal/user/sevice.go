package user

import (
	"context"
	"log"
	pb "rpcprac/pb/user/proto"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	Storage *UserStorage
}

func (s *UserService) CreateUser(
	ctx context.Context,
	req *pb.CreateUserRequest,
) (*pb.CreateUserResponse, error) {

	log.Println("CreateUser called with name:", req.Name)
	user := s.Storage.CreateUser(req.Name)

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Level: user.Level,
			Xp:    user.Xp,
		},
	}, nil
}
func (s *UserService) GetUser(
	ctx context.Context,
	req *pb.GetUserRequest,
) (*pb.GetUserResponse, error) {

	log.Println("GetUser called with id:", req.Id)
	user := s.Storage.GetUser(req.Id)
	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Level: user.Level,
			Xp:    user.Xp,
		},
	}, nil
}
func (s *UserService) ListUsers(
	ctx context.Context,
	req *pb.ListUsersRequest,
) (*pb.ListUsersResponse, error) {

	users := s.Storage.ListUsers()

	pbUsers := make([]*pb.User, 0, len(users))
	for _, user := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Level: user.Level,
			Xp:    user.Xp,
		})
	}

	log.Printf("returning %d users", len(pbUsers))

	return &pb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
