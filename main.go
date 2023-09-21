package main

import (
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmd/root"
	"github.com/tradlwa/xj/cmdutil"
)

func main() {
	client := api.NewClient("http://192.168.201.3:8084")
	context := &cmdutil.Context{ApiClient: client}
	cmdRoot := root.NewCmdRoot(context)
	_ = cmdRoot.Execute()
}
