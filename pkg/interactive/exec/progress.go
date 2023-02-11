package exec

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

var depth = 0

type Progress struct {
	string
}

func Start(p any) *Progress {
	depth++
	pterm.Info.Println(strings.Repeat("+ ", depth), p)
	return &Progress{fmt.Sprint(p)}
}

func (p Progress) Stop(err error) {
	defer func() { depth-- }()
	if err != nil {
		pterm.Error.Println(strings.Repeat("+ ", depth), p.string)
		return
	}
	pterm.Success.Println(strings.Repeat("+ ", depth), p.string)
}
