package rpc

import (
	"context"
	pb "github.com/theshadow/cmdnctrl/echo"
	"google.golang.org/grpc"
)

type Server struct{
	rpcSrv *grpc.Server
	done   chan struct{}
}

func New(rpc *grpc.Server, done chan struct{}) *Server {
	return &Server{rpcSrv: rpc, done: done}
}

func (s *Server) Echo(ctx context.Context, in *pb.EchoRequest) (*pb.EchoReply, error) {
	return &pb.EchoReply{Message: in.Message}, nil
}

func (s *Server) Shutdown(ctx context.Context, in *pb.ShutdownRequest) (*pb.ShutdownResponse, error) {
	close(s.done)
	return &pb.ShutdownResponse{}, nil
}


