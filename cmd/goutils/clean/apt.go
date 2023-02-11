package clean

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var APTCmd = &cobra.Command{
	Use: "apt",
	Run: func(cmd *cobra.Command, args []string) {
		c := exec.Command("sudo", "apt", "autoremove")
		_errors.Check(c.InteractiveRun())
	},
}
