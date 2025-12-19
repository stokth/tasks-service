package grpc

import (
	"net"

	taskpb "github.com/stokth/project-protos/proto/task"
	userpb "github.com/stokth/project-protos/proto/user"
	"github.com/stokth/tasks-service/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc *task.Service, uc userpb.UserServiceClient) error {
	lis, _ := net.Listen("tcp", ":50052")
	grpcSrv := grpc.NewServer()
	handler := NewHandler(svc, uc)
	taskpb.RegisterTaskServiceServer(grpcSrv, handler)
	return grpcSrv.Serve(lis)
}
