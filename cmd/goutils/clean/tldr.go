package clean

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var TLDRCmd = &cobra.Command{
	Use: "tldr",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("tldr", "--clean-cache")
		_errors.Check(c.InteractiveRun())
	},
}
