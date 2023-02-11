package _import

import (
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "import",
	Run: exec.RunCommands,
}

func init() {
	RootCmd.AddCommand(GPGCmd)
}
