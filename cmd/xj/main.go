package main

import (
	"fmt"

	"github.com/heminghu/xj/internal/cmd/root"
	"github.com/heminghu/xj/internal/cmdcontext"
	"github.com/heminghu/xj/internal/config"
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
