package update

import (
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "update",
	Run: exec.RunCommands,
}

func init() {
	RootCmd.AddCommand(APTCmd)
	RootCmd.AddCommand(BrewCmd)
	RootCmd.AddCommand(SnapCmd)
	RootCmd.AddCommand(TLDRCmd)
	RootCmd.AddCommand(NpmCmd)
}
