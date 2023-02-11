package export

import (
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "export",
	Run: exec.RunCommands,
}

func init() {
	RootCmd.AddCommand(GPGCmd)
}
