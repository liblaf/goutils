package sort

import (
	"github.com/liblaf/goutils/cmd/goutils/sort/awesome"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "sort",
}

func init() {
	RootCmd.AddCommand(awesome.RootCmd)
}
