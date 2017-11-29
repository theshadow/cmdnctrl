package cmd

import (
	"log"
	"net"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/theshadow/cmdnctrl/rpc"
	pb "github.com/theshadow/cmdnctrl/echo"
)

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the simple echo gRPC service",
	Long: `Starts the simple echo gRPC service`,
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		done := make(chan struct{})

		srv := grpc.NewServer()
		pb.RegisterEchoerServer(srv, rpc.New(srv, done))
		reflection.Register(srv)

		go srv.Serve(lis)

		<-done

		srv.GracefulStop()
	},
}


