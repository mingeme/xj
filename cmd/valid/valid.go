package valid

import (
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
)

var mustArgsRe = regexp.MustCompile("<(.*?)>")

func ExpectedArgs() cobra.PositionalArgs {
	return func(cmd *cobra.Command, args []string) error {
		requiredArgs := mustArgsRe.FindAllString(cmd.Use, -1)
		if len(requiredArgs) <= len(args) {
			return nil
		}
		missingArgs := requiredArgs[len(args):]
		joined := strings.Join(missingArgs, ", ")
		return fmt.Errorf(`required argument(s) "%s" not set`, joined)
	}
}
