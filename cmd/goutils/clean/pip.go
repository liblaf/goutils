package clean

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var PipCmd = &cobra.Command{
	Use: "pip",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("conda", "clean", "--all")
		_errors.Check(c.InteractiveRun())
		c = exec.Command("pip", "cache", "purge")
		_errors.Check(c.InteractiveRun())
	},
}
