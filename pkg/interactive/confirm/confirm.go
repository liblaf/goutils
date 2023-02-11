package confirm

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/message"
	_path "github.com/liblaf/goutils/pkg/path"
	"github.com/pterm/pterm"
)

var Yes bool = true

func Confirm(msg *message.Message, defaultValue bool) bool {
	if Yes {
		return true
	}

	rsp, err := pterm.DefaultInteractiveConfirm.WithDefaultValue(defaultValue).Show(msg.String())
	_errors.Check(err)
	return rsp
}

func ConfirmOnExists(msg *message.Message, defaultValue bool) bool {
	exists, err := _path.Exists(msg.Destination)
	_errors.Check(err)
	if exists {
		return Confirm(msg, defaultValue)
	} else {
		return true
	}
}
