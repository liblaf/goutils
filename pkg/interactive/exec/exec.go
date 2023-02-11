package exec

import (
	"os"
	"os/exec"
	"strings"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/pterm/pterm"
	"mvdan.cc/sh/v3/syntax"
)

var cmdStyle = pterm.ThemeDefault.InfoMessageStyle

type Cmd struct {
	*exec.Cmd
}

func Command(name string, arg ...string) *Cmd {
	return &Cmd{exec.Command(name, arg...)}
}

func (c *Cmd) String() string {
	var args []string
	for _, arg := range c.Args {
		arg, err := syntax.Quote(arg, syntax.LangBash)
		_errors.Check(err)
		args = append(args, arg)
	}
	return strings.Join(args, " ")
}

func (c *Cmd) InteractiveRun() error {
	p := Start(c)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Cmd.Run()
	p.Stop(err)
	return err
}

func (c *Cmd) Run() error {
	depth++
	defer func() { depth-- }()
	spinner, err := pterm.DefaultSpinner.WithMessageStyle(&cmdStyle).Start(strings.Repeat("+ ", depth), c)
	_errors.Check(err)
	output, err := c.Cmd.CombinedOutput()
	if err != nil {
		spinner.Fail()
		pterm.Error.Println(string(output))
		return err
	}
	spinner.Success()
	return nil
}
