package main

import (
	"github.com/liblaf/goutils/cmd/goutils/assets"
	"github.com/liblaf/goutils/cmd/goutils/clean"
	"github.com/liblaf/goutils/cmd/goutils/keys"
	"github.com/liblaf/goutils/cmd/goutils/update"
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/confirm"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "goutils",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		yes, err := cmd.Flags().GetBool("yes")
		_errors.Check(err)
		confirm.Yes = yes
	},
}

func init() {
	RootCmd.PersistentFlags().BoolP("yes", "y", false, "")

	RootCmd.AddCommand(assets.RootCmd)
	RootCmd.AddCommand(keys.RootCmd)
	RootCmd.AddCommand(clean.RootCmd)
	RootCmd.AddCommand(update.RootCmd)
}
