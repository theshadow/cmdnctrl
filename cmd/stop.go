package cmd

import (
	"context"
	"fmt"
	"time"
	"github.com/spf13/cobra"
	"github.com/theshadow/cmdnctrl/echo"
	"google.golang.org/grpc"
)

var StopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stops the simple echo gRPC service",
	Long: `Stops the simple echo gRPC service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 6000)
		defer cancel()

		conn, err := grpc.DialContext(ctx, "127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("unable to dial service %s", err)
		}

		c := echo.NewEchoerClient(conn)
		_, err = c.Shutdown(ctx, &echo.ShutdownRequest{})
		if err != nil {
			return fmt.Errorf("unable to execute Shutdown()! %s", err)
		}

		return nil
	},
}
