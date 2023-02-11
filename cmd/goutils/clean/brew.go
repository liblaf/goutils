package clean

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var BrewCmd = &cobra.Command{
	Use: "brew",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("brew", "autoremove")
		_errors.Check(c.InteractiveRun())
		c = exec.Command("brew", "cleanup")
		_errors.Check(c.InteractiveRun())
	},
}
