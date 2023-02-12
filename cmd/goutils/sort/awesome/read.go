package awesome

import (
	"os"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"gopkg.in/yaml.v3"
)

func read(filename string) awesomeList {
	raw, err := os.ReadFile(filename)
	_errors.Check(err)

	list := make(awesomeList)
	yaml.Unmarshal(raw, list)

	return list
}
