package tool

import (
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"testing"
)

func TestGet(t *testing.T) {
	url := "https://cdn152.akamai-content-network.com/bcdn_token=1ZgG3_nX6r2qxtdDZT2GpQa6XSzfa-vSBgaeNumX9F8&expires=1692063155&token_path=%2F3bf3f2d9-7470-4744-8ffe-ec5832b425af%2F/3bf3f2d9-7470-4744-8ffe-ec5832b425af/842x480/video8.ts"
	referer := "https://666.com/"
	body, err := Get(url, referer)
	if err != nil {
		t.Error(err)
	}
	defer body.Close()

	data, err := ioutil.ReadAll(body)
	if err != nil {
		t.Error(err)
	}
	spew.Dump(string(data))
}
