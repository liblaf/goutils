package _import

import (
	"path"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/spf13/cobra"
)

var GPGCmd = &cobra.Command{
	Use: "gpg",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := cmd.Flags().GetString("path")
		_errors.Check(err)
		c := exec.Command("gpg", "--import", path.Join(p, "gpg", "secret.asc"))
		_errors.Check(c.InteractiveRun())
	},
}
