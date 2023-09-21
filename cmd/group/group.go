package group

import (
	"bytes"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/cmdutil"
)

func NewCmdGroup(c *cmdutil.Context) *cobra.Command {
	groupCmd := &cobra.Command{
		Use:     "group",
		Aliases: []string{"g"},
		Short:   "Management for job group",
	}

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "执行器查询",
		Run: func(cmd *cobra.Command, args []string) {
			_ = c.ApiClient.Post("jobgroup/pageList", bytes.NewBufferString("appname=&title=&start=0&length=10"), nil)
		},
	}

	groupCmd.AddCommand(lsCmd)
	return groupCmd
}
