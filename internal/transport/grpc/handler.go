package grpc

import (
	"context"
	"fmt"

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

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	// 1. Проверить пользователя:
	if _, err := h.userClient.GetUser(ctx, &userpb.GetUserRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}
	// 2. Внутренняя логика:
	t, err := h.svc.CreateTask(&task.Task{UserID: req.UserId, Title: req.Title})
	if err != nil {
		return nil, err
	}
	// 3. Ответ:
	return &taskpb.CreateTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Title}}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *taskpb.ListTasksRequest) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasks()
	if err != nil {
		return nil, err
	}

	var pbTasks []*taskpb.Task
	for _, t := range tasks {
		pbTasks = append(pbTasks, &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Title})
	}

	return &taskpb.ListTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	t, err := h.svc.GetTask(req.Id)
	if err != nil {
		return nil, err
	}

	return &taskpb.GetTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Title}}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	t, err := h.svc.UpdateTask(req.Id, &task.Task{Title: req.Title})
	if err != nil {
		return nil, err
	}

	return &taskpb.UpdateTaskResponse{Task: &taskpb.Task{Id: t.ID, UserId: t.UserID, Title: t.Title}}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	if err := h.svc.DeleteTask(req.Id); err != nil {
		return nil, err
	}

	return &taskpb.DeleteTaskResponse{}, nil
}
