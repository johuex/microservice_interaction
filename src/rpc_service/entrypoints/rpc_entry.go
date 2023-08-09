package entrypoints

import (
	"context"
	pb "rpc_compiled"
	"rpc_service/shared"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedExampleServiceServer
}

func (c *Server) Call(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	res := shared.ContainerItem.Service.Call(int(in.Message))
	out := pb.Response{Message: int32(res)}
	return &out, nil
}

func NewGRPCServer() *grpc.Server {
	gsrv := grpc.NewServer()
	pb.RegisterExampleServiceServer(gsrv, &Server{})
	return gsrv
}
