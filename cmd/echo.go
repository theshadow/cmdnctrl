package cmd

import (
	"context"
	"fmt"
	"time"
	"github.com/spf13/cobra"
	"github.com/theshadow/cmdnctrl/echo"
	"google.golang.org/grpc"
)

var EchoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Calls the Echo() method via RPC and prints the results",
	Long: `Calls the Echo() method via RPC and prints the results`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing argument")
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond * 6000)
		defer cancel()

		conn, err := grpc.DialContext(ctx, "127.0.0.1:50051", grpc.WithInsecure())
		if err != nil {
			return fmt.Errorf("unable to dial service %s", err)
		}

		c := echo.NewEchoerClient(conn)
		resp, err := c.Echo(ctx, &echo.EchoRequest{Message: args[0]})
		if err != nil {
			return fmt.Errorf("unable to execute Echo()! %s", err)
		}

		fmt.Println(resp.Message)

		return nil
	},
}
