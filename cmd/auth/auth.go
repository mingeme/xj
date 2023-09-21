package auth

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tradlwa/xj/cmd/valid"
	"github.com/tradlwa/xj/cmdutil"
)

func NewCmdAuth(c *cmdutil.Context) *cobra.Command {
	command := &cobra.Command{
		Use:   "auth <domain> <username>",
		Short: "Authenticate for domain",
		Args:  valid.ExpectedArgs(),
		Run: func(cmd *cobra.Command, args []string) {
			domain := args[0]
			username := args[1]
			fmt.Printf("domain=%s username=%s", domain, username)
		},
	}
	return command
}
