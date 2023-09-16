package tool

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Get(url string, referer string) (io.ReadCloser, error) {
	var c http.Client
	tr := GetProxyTr(url)
	if tr != nil {
		c = http.Client{
			Timeout:   time.Duration(15) * time.Second, // timeout
			Transport: GetProxyTr(url),
		}
	} else {
		c = http.Client{
			Timeout: time.Duration(15) * time.Second, // timeout
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/117.0")
	if referer != "" {
		req.Header.Add("Referer", referer)
	}

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

func GetProxyTr(resUrl string) (tr *http.Transport) {
	if strings.Contains(resUrl, "127.0.0.1") || strings.Contains(resUrl, "127.0.0.1") {
		return nil
	}
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
