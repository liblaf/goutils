package main

import (
	er "github.com/liblaf/goutils/pkg/errors"
)

func init() {
	er.DefaultLogger = er.PTerm
}

func main() {
	er.Check(RootCmd.Execute())
}
