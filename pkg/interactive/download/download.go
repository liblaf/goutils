package download

import (
	"io"
	"net/http"
	"os"
	"sync"

	_errors "github.com/liblaf/goutils/pkg/errors"
	"github.com/liblaf/goutils/pkg/interactive/confirm"
	"github.com/liblaf/goutils/pkg/message"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

func doDownload(src string, dst string) {
	var wg sync.WaitGroup
	p := mpb.New(mpb.WithWaitGroup(&wg))

	req, err := http.NewRequest(http.MethodGet, src, nil)
	_errors.Check(err)
	req.Header.Set("User-Agent", "Mozilla")

	wg.Add(1)
	go func() {
		defer wg.Done()

		resp, err := http.DefaultClient.Do(req)
		_errors.Check(err)
		defer resp.Body.Close()

		b := p.AddBar(resp.ContentLength,
			mpb.PrependDecorators(
				decor.Name(dst, decor.WCSyncSpaceR),
				decor.Percentage(decor.WCSyncSpaceR),
				decor.CountersKibiByte("(%.1f/%.1f)", decor.WCSyncSpaceR),
				decor.Elapsed(decor.ET_STYLE_MMSS, decor.WCSyncSpaceR),
				decor.OnComplete(
					decor.AverageETA(decor.ET_STYLE_MMSS, decor.WCSyncSpaceR),
					"Done.",
				),
				decor.AverageSpeed(decor.UnitKiB, "% .2f", decor.WCSyncSpaceR),
			),
		)
		r := b.ProxyReader(resp.Body)
		defer r.Close()

		out, err := os.Create(dst)
		_errors.Check(err)
		defer out.Close()

		n, err := io.Copy(out, r)
		_errors.Check(err)

		if resp.ContentLength < 0 {
			b.SetTotal(n, true)
		}
	}()

	p.Wait()
}

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

	doDownload(src, dst)

	return true
}
