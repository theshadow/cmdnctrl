package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "cmdnctrl",
	Short: "CmdNCtrl is a small Asynchronous gRPC proof of concept.",
	Long: `CmdNCtrl is a small Asynchronous gRPC proof of concept.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(EchoCmd)
	RootCmd.AddCommand(StartCmd)
	RootCmd.AddCommand(StopCmd)
}
