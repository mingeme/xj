package root

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/cmd/auth"
	"github.com/tradlwa/xj/cmd/group"
	"github.com/tradlwa/xj/cmd/job"
	"github.com/tradlwa/xj/cmdcontext"
)

func NewCmdRoot(c *cmdcontext.Context) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "xj",
		Short: "xxl-job CLI",
	}

	cmd.PersistentFlags().StringVarP(&c.Env, "env", "e", "", "an alias for domain")
	err := cmd.MarkPersistentFlagRequired("env")
	if err != nil {
		fmt.Println(err)
	}

	cmd.AddCommand(auth.NewCmdAuth(c))
	cmd.AddCommand(group.NewCmdGroup(c))
	cmd.AddCommand(job.NewCmdTask(c))

	return cmd
}
