package job

import (
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/cmdutil"
)

func NewCmdTask(c *cmdutil.Context) *cobra.Command {
	taskCmd := &cobra.Command{
		Use:     "job",
		Aliases: []string{"j"},
		Short:   "Management for job",
	}

	lsCmd := &cobra.Command{
		Use:   "ls",
		Short: "任务查询",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	taskCmd.AddCommand(lsCmd)
	return taskCmd
}
