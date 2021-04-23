package tool

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func Get(url string) (io.ReadCloser, error) {
	c := http.Client{
		Timeout: time.Duration(60) * time.Second,
	}
	// if strings.HasSuffix(url, ".ts") {
	// 	url = strings.Replace(url, "ts3.510yh.cc/", "ts3.510yh.cc:4439/", -1)
	// 	fmt.Println("18", url)
	// } else {
	// 	fmt.Println("21", url)
	// }
	resp, err := c.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http error: status code %d", resp.StatusCode)
	}
	return resp.Body, nil
}
