package assets

import (
	"github.com/liblaf/goutils/cmd/goutils/assets/avatar"
	"github.com/liblaf/goutils/cmd/goutils/assets/favicon"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "assets",
	Run: exec.RunCommands,
}

func init() {
	RootCmd.AddCommand(avatar.RootCmd)
	RootCmd.AddCommand(favicon.RootCmd)
}
