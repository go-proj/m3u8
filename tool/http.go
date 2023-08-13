package tool

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string, referer string) (io.ReadCloser, error) {
	c := http.Client{
		Timeout: time.Duration(3) * time.Second, // timeout
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	req.Header.Add("Referer", referer)
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
