package export

import (
	"path"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	_os "github.com/liblaf/goutils/pkg/os"
	"github.com/spf13/cobra"
)

var GPGCmd = &cobra.Command{
	Use: "gpg",
	Run: func(cmd *cobra.Command, args []string) {
		p, err := cmd.Flags().GetString("path")
		_errors.Check(err)
		_errors.Check(_os.Mkdir(path.Join(p, "gpg")))
		c := exec.Command("gpg", "--export-secret-keys", "--armor", "--output", path.Join(p, "gpg", "secret.asc"))
		_errors.Check(c.InteractiveRun())
	},
}
