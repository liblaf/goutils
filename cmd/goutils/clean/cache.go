package clean

import (
	"os"
	"path/filepath"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/remove"
	"github.com/spf13/cobra"
)

var patterns []string

func init() {
	cache, err := os.UserCacheDir()
	_errors.Check(err)
	home, err := os.UserHomeDir()
	_errors.Check(err)
	patterns = append(patterns,
		cache,
		filepath.Join(home, ".profile"),
		filepath.Join(home, "*bash*"),
		filepath.Join(home, "*zcompdump*"),
	)
}

func GlobPathes(patterns []string) (rtn []string) {
	for _, p := range patterns {
		newPathes, err := filepath.Glob(p)
		_errors.Check(err)
		rtn = append(rtn, newPathes...)
	}
	return
}

var CacheCmd = &cobra.Command{
	Use: "cache",
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range GlobPathes(patterns) {
			remove.Remove(p)
		}
	},
}
