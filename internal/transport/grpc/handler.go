package grpc

import (
	taskpb "github.com/stokth/project-protos/proto/task"
	userpb "github.com/stokth/project-protos/proto/user"
	"github.com/stokth/tasks-service/internal/task"
)

type Handler struct {
	svc        *task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(svc *task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}
