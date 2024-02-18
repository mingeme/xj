package interactive

import "github.com/spf13/cobra"

func NewCmdInteractive() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interactive",
		Aliases: []string{"i"},
		Short:   "Enter interactive mode",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	return cmd
}
