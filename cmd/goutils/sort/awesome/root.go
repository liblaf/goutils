package awesome

import (
	"context"
	"os"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var RootCmd = &cobra.Command{
	Use: "awesome",

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		output, err := cmd.Flags().GetString("output")
		_errors.Check(err)
		token, err := cmd.Flags().GetString("token")
		_errors.Check(err)

		list := read(file)
		initTemplate()
		s := sortAll(context.Background(), list, token)
		y, err := yaml.Marshal(list)
		_errors.Check(err)
		_errors.Check(os.WriteFile(file, y, 0o644))
		_errors.Check(os.WriteFile(output, []byte(s), 0o644))
	},
}

func init() {
	RootCmd.Flags().StringP("output", "o", "awesome.md", "")
	RootCmd.MarkFlagFilename("output")
	RootCmd.Flags().StringP("token", "t", "", "")
}
