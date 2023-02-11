package clean

import (
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "clean",
	Run: exec.RunCommands,
}

func init() {
	RootCmd.AddCommand(APTCmd)
	RootCmd.AddCommand(BrewCmd)
	RootCmd.AddCommand(CacheCmd)
	RootCmd.AddCommand(NpmCmd)
	RootCmd.AddCommand(PipCmd)
	RootCmd.AddCommand(TLDRCmd)
	RootCmd.AddCommand(TmpCmd)
}
