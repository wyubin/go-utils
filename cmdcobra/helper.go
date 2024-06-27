package cmdcobra

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GetScript(ccmd *cobra.Command) string {
	flags := ccmd.Flags()
	argsPos := strings.Join(flags.Args(), " ")
	argsOther := []string{}
	flags.VisitAll(func(flag *pflag.Flag) {
		if !flag.Changed {
			return
		}
		argsOther = append(argsOther, fmt.Sprint("--", flag.Name))
		argsOther = append(argsOther, flag.Value.String())
	})
	return fmt.Sprintf("%s %s %s", ccmd.CommandPath(), argsPos, strings.Join(argsOther, " "))
}
