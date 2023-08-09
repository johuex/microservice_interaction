package main

import (
	"net"
	"rpc_service/entrypoints"

	"google.golang.org/grpc/grpclog"
)

func main() {
	listener, err := net.Listen("tcp", ":9090")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	srv := entrypoints.NewGRPCServer()
	grpclog.Info("Serve on :9090")
	srv.Serve(listener)
}
