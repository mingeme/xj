package valid

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
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

func RequireInt(arg string) (int, error) {
	if val, err := strconv.Atoi(arg); err == nil {
		return val, nil
	}
	return 0, fmt.Errorf("require int argument: %s", arg)
}
