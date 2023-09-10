package tool

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func Get(url string, referer string) (io.ReadCloser, error) {
	c := http.Client{
		Timeout:   time.Duration(9) * time.Second, // timeout
		Transport: GetProxyTr(),
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	req.Header.Add("Referer", referer)

	var resp *http.Response
	for i := 0; i < 3; i++ {
		resp, err = c.Do(req)
		if err != nil {
			fmt.Printf("[error] read %s [%+v]", url, err)
			continue // retry
		}
		//defer resp.Body.Close() //❌这里不能 close，后面要用
	}
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}

func GetProxyTr() (tr *http.Transport) {
	proxyUrl := os.Getenv("http_proxy")
	if proxyUrl == "" {
		return nil
	}

	tr = &http.Transport{
		Proxy: func(*http.Request) (*url.URL, error) {
			return url.Parse(proxyUrl)
		},
	}
	return tr
}
