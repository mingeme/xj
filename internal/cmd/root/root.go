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

	cmd.AddCommand(envFlagRequired(auth.NewCmdAuth(c), err))
	cmd.AddCommand(envFlagRequired(group.NewCmdGroup(c), err))
	cmd.AddCommand(envFlagRequired(job.NewCmdJob(c), err))

	return cmd
}

func envFlagRequired(cmd *cobra.Command, err error) *cobra.Command {
	if len(cmd.Commands()) > 0 {
		for _, subCommand := range cmd.Commands() {
			envFlagRequired(subCommand, err)
		}
	} else {
		cmd.Run = func(cmd *cobra.Command, args []string) {
			if err != nil {
				fmt.Println(err)
				return
			}
			cmd.Run(cmd, args)
		}
	}

	return cmd
}
