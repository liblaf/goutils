package errors

import (
	"github.com/pterm/pterm"
	"github.com/sirupsen/logrus"
)

type Logger int

const (
	Logrus Logger = iota
	PTerm
)

var DefaultLogger Logger = Logrus

func Fatalln(a ...any) {
	switch DefaultLogger {
	case Logrus:
		logrus.Fatalln(a...)
	case PTerm:
		pterm.Fatal.Println(a...)
	}
}

func Check(err error) bool {
	if err != nil {
		Fatalln(err)
		return false
	}

	return true
}
