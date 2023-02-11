package os

import (
	"os"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/message"
	"github.com/pterm/pterm"
)

func Mkdir(p string) error {
	msg := message.Message{
		Action: "MKDIR",
		Source: p,
	}
	spinner, err := pterm.DefaultSpinner.Start(msg)
	_errors.Check(err)
	err = os.MkdirAll(p, os.ModePerm)
	if err != nil {
		spinner.Fail(err)
		return err
	}
	spinner.Success()
	return nil
}
