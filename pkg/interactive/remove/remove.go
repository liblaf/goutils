package remove

import (
	"errors"
	"io/fs"
	"os"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/confirm"
	"github.com/liblaf/goutils/pkg/message"
	"github.com/pterm/pterm"
)

func Remove(p string) bool {
	msg := &message.Message{
		Action:      "REMOVE",
		Destination: p,
	}
	rsp := confirm.ConfirmOnExists(msg, true)
	if !rsp {
		return false
	}

	spinner, err := pterm.DefaultSpinner.Start(msg)
	_errors.Check(err)
	err = os.RemoveAll(p)
	if err != nil {
		switch {
		case errors.Is(err, fs.ErrPermission):
		default:
			pterm.Error.Println(err)
		}
		spinner.RemoveWhenDone = true
		spinner.Fail()
	}
	spinner.Success()
	return true
}
