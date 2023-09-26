package job

import (
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/api"
	"github.com/tradlwa/xj/cmdcontext"
)

func NewCmdTask(c *cmdcontext.Context) *cobra.Command {
	taskCmd := &cobra.Command{
		Use:     "job",
		Aliases: []string{"j"},
		Short:   "Management for job",
	}

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "任务查询",
		Run: func(cmd *cobra.Command, args []string) {
			api.JobPage(c.ApiClient(), &api.JobOptions{Start: 0, Length: 10})
		},
	}

	taskCmd.AddCommand(lsCmd)
	return taskCmd
}
