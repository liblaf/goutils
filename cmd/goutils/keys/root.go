package keys

import (
	"github.com/liblaf/goutils/cmd/goutils/keys/_import"
	"github.com/liblaf/goutils/cmd/goutils/keys/export"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "keys",
}

func init() {
	RootCmd.PersistentFlags().StringP("path", "p", "keys", "")

	RootCmd.AddCommand(export.RootCmd)
	RootCmd.AddCommand(_import.RootCmd)
}
