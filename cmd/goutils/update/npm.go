package update

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var NpmCmd = &cobra.Command{
	Use: "npm",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("pnpm", "update", "--global")
		_errors.Check(c.InteractiveRun())
	},
}
