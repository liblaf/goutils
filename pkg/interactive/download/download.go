package download

import (
	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/confirm"
	"github.com/liblaf/goutils/pkg/interactive/exec"
	"github.com/liblaf/goutils/pkg/message"
)

func Download(src string, dst string) bool {
	msg := &message.Message{
		Action:      "DOWNLOAD",
		Source:      src,
		Destination: dst,
	}
	rsp := confirm.ConfirmOnExists(msg, true)
	if !rsp {
		return false
	}
	c := exec.Command("wget", "--output-document", dst, src)
	_errors.Check(c.Run())
	return true
}
