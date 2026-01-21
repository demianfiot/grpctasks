package handler

import (
	"context"

	pb "rpcprac/pb/user/proto"
	"rpcprac/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService service.User
}

func NewUserHandler(userService service.User) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := h.userService.CreateUser(req.Name)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to create user: %v",
			err,
		)
	}

	return &pb.CreateUserResponse{
		User: &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Level: user.Level,
			Xp:    user.Xp,
		},
	}, nil
}

func (h *UserHandler) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {

	user, err := h.userService.GetUser(req.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"user with id %d not found",
			req.Id,
		)
	}

	return &pb.GetUserResponse{
		User: &pb.User{
			Id:    user.Id,
			Name:  user.Name,
			Level: user.Level,
			Xp:    user.Xp,
		},
	}, nil
}

func (h *UserHandler) ListUsers(ctx context.Context, _ *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {

	users, err := h.userService.ListUsers()
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"failed to list users: %v",
			err,
		)
	}

	pbUsers := make([]*pb.User, 0, len(users))
	for _, u := range users {
		pbUsers = append(pbUsers, &pb.User{
			Id:    u.Id,
			Name:  u.Name,
			Level: u.Level,
			Xp:    u.Xp,
		})
	}

	return &pb.ListUsersResponse{
		Users: pbUsers,
	}, nil
}
