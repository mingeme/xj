package root

import (
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/cmd/auth"
	"github.com/tradlwa/xj/cmd/group"
	"github.com/tradlwa/xj/cmd/job"
	"github.com/tradlwa/xj/cmdutil"
)

func NewCmdRoot(c *cmdutil.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xj",
		Short: "xxl-job CLI",
	}

	cmd.AddCommand(auth.NewCmdAuth(c))
	cmd.AddCommand(group.NewCmdGroup(c))
	cmd.AddCommand(job.NewCmdTask(c))

	return cmd
}
