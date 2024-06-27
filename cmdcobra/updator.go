package cmdcobra

import (
	"github.com/spf13/cobra"
)

type CmdUpdator func(ccmd *cobra.Command) error

func CmdUpdatorVer(strVer string) CmdUpdator {
	return func(ccmd *cobra.Command) error {
		ccmd.Version = strVer
		return nil
	}
}
