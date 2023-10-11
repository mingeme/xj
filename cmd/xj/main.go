package main

import (
	"fmt"

	"github.com/tradlwa/xj/internal/cmd/root"
	"github.com/tradlwa/xj/internal/cmdcontext"
	"github.com/tradlwa/xj/internal/config"
)

func main() {
	envConfig, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("cannot read config file %v", err)
		return
	}
	context := &cmdcontext.Context{EnvConfig: envConfig}
	cmdRoot := root.NewCmdRoot(context)
	_ = cmdRoot.Execute()
}
