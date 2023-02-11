package message

import (
	"fmt"
	"path"
	"regexp"

	_errors "github.com/liblaf/goutils/pkg/errors"
	_path "github.com/liblaf/goutils/pkg/path"
	"github.com/pterm/pterm"
)

// https://regex101.com/r/cX0pJ8/1
var urlRegex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[.\!\/\\w]*))?)`)

var (
	actionStyle   = pterm.NewStyle(pterm.Bold, pterm.FgMagenta)
	filenameStyle = pterm.NewStyle(pterm.FgMagenta)
	pathStyle     = pterm.NewStyle(pterm.FgLightMagenta)
	urlStyle      = pterm.NewStyle(pterm.Underscore, pterm.FgLightBlue)
)

var (
	to        = pterm.Bold.Sprint("->")
	overwrite = pterm.ThemeDefault.WarningMessageStyle.Sprint("(OVERWRITE)")
)

func prettyPath(s string) string {
	dir, file := path.Split(s)
	return pathStyle.Sprint(dir) + filenameStyle.Sprint(file)
}

func prettyURL(s string) string {
	return urlStyle.Sprint(s)
}

func pretty(s string) string {
	switch {
	case urlRegex.MatchString(s):
		return prettyURL(s)
	default:
		return prettyPath(s)
	}
}

func (msg Message) String() string {
	var action, src, dst string
	var err error

	if msg.Action != "" {
		action = actionStyle.Sprint(msg.Action)
	}
	if msg.Source != "" {
		src = pretty(msg.Source)
		_errors.Check(err)
	}
	if msg.Destination != "" {
		exists, err := _path.Exists(msg.Destination)
		_errors.Check(err)
		if exists {
			dst = pretty(msg.Destination)
			_errors.Check(err)
			dst += " " + overwrite
		} else {
			dst = pretty(msg.Destination)
			_errors.Check(err)
		}
	}

	if action != "" {
		switch {
		case src != "" && dst != "":
			return fmt.Sprintf("%s %s %s %s", action, src, to, dst)
		case src != "" && dst == "":
			return fmt.Sprintf("%s %s", action, src)
		case src == "" && dst != "":
			return fmt.Sprintf("%s %s", action, dst)
		default:
			return action
		}
	} else {
		switch {
		case src != "" && dst != "":
			return fmt.Sprintf("%s %s %s", src, to, dst)
		case src != "" && dst == "":
			return src
		case src == "" && dst != "":
			return dst
		default:
			return ""
		}
	}
}
