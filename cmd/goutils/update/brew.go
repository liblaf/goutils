package update

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var BrewCmd = &cobra.Command{
	Use: "brew",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("brew", "update")
		_errors.Check(c.InteractiveRun())
		c = exec.Command("brew", "upgrade")
		_errors.Check(c.InteractiveRun())
	},
}
