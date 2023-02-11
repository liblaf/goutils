package update

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var SnapCmd = &cobra.Command{
	Use: "snap",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("sudo", "snap", "refresh")
		_errors.Check(c.InteractiveRun())
	},
}
