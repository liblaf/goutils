package clean

import (
	"path/filepath"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/remove"
	"github.com/spf13/cobra"
)

var TmpCmd = &cobra.Command{
	Use: "tmp",
	Run: func(cmd *cobra.Command, args []string) {
		pathes, err := filepath.Glob("/tmp/*")
		_errors.Check(err)
		for _, p := range pathes {
			remove.Remove(p)
		}
	},
}
