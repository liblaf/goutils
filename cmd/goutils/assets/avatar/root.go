package avatar

import (
	"fmt"
	"path/filepath"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/download"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/liblaf/goutils/pkg/os"
	"github.com/spf13/cobra"
)

var formats = []string{"jpg", "png"}

var output string

var RootCmd = &cobra.Command{
	Use: "avatar",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		output, err = cmd.Flags().GetString("output")
		_errors.Check(err)
	},

	Run: func(cmd *cobra.Command, args []string) {
		for _, format := range formats {
			os.Mkdir(filepath.Join(output, format))
		}
		for name, url := range urls {
			jpg := filepath.Join(output, "jpg", fmt.Sprintf("%s.jpg", name))
			png := filepath.Join(output, "png", fmt.Sprintf("%s.png", name))
			download.Download(url, jpg)
			_errors.Check(exec.Command("magick", "convert", jpg, png).Run())
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringP("output", "o", "avatar", "")
	RootCmd.MarkPersistentFlagDirname("output")
}
