package favicon

import (
	"fmt"
	"path/filepath"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/download"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/liblaf/goutils/pkg/os"
	"github.com/spf13/cobra"
)

var formats = []string{"svg", "png", "ico"}

var color, output string

var RootCmd = &cobra.Command{
	Use: "favicon",

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		color, err = cmd.Flags().GetString("color")
		_errors.Check(err)
		output, err = cmd.Flags().GetString("output")
		_errors.Check(err)
	},

	Run: func(cmd *cobra.Command, args []string) {
		for _, format := range formats {
			os.Mkdir(filepath.Join(output, format))
		}
		for letter := 'a'; letter <= 'z'; letter++ {
			ico := filepath.Join(output, "ico", fmt.Sprintf("%c.ico", letter))
			png := filepath.Join(output, "png", fmt.Sprintf("%c.png", letter))
			svg := filepath.Join(output, "svg", fmt.Sprintf("%c.svg", letter))
			download.Download(fmt.Sprintf("https://raw.githubusercontent.com/FortAwesome/Font-Awesome/6.x/svgs/solid/%c.svg", letter), svg)
			args := []string{"magick",
				"convert",
				"-background",
				"none",
				svg,
				"-fill",
				color,
				"-colorize",
				"100",
				"-resize",
				"512x512",
				"-gravity",
				"center",
				"-extent",
				"512x512",
			}
			_errors.Check(exec.Command(args[0], append(args[1:], png)...).Run())
			_errors.Check(exec.Command(args[0], append(args[1:], "-resize", "128x128", ico)...).Run())
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringP("color", "c", "#48BEF3", "")
	RootCmd.PersistentFlags().StringP("output", "o", "favicon", "")
}
