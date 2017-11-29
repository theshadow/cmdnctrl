
//go:generate protoc -I echo/ echo/echo.proto --go_out=plugins=grpc:echo

package main

import (
	"github.com/theshadow/cmdnctrl/cmd"
)

func main() {
	cmd.RootCmd.Execute()
}
