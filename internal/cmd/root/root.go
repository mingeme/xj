package root

import (
	"fmt"

	"github.com/heminghu/xj/internal/cmd/auth"
	"github.com/heminghu/xj/internal/cmd/group"
	"github.com/heminghu/xj/internal/cmd/job"
	"github.com/heminghu/xj/internal/cmdcontext"
	"github.com/spf13/cobra"
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
	cmd.AddCommand(job.NewCmdJob(c))

	return cmd
}
