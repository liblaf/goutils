package exec

import (
	"strings"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func RunCommands(cmd *cobra.Command, args []string) {
	root := cmd.Root()
	for _, c := range cmd.Commands() {
		pterm.DefaultSection.Println(c.CommandPath())
		root.SetArgs(append(strings.Fields(c.CommandPath())[1:], args...))
		_errors.Check(root.Execute())
	}
}
