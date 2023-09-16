package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"coding.net/cy9/golang/m3u8/dl"
)

var (
	url      string
	output   string
	referer  string
	chanSize int
	tsPort   int
)

// 实际遇到过 key 和 ts 文件，端口可能跟 m3u8 文件不同
func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required")
	flag.StringVar(&output, "o", "", "Output folder, required")
	flag.IntVar(&chanSize, "c", 32, "Maximum number of occurrences")

	flag.IntVar(&tsPort, "t", 0, "实际遇到过 key 和 ts 文件，端口可能跟 m3u8 文件不同")
	flag.StringVar(&referer, "r", "", "部分 ts 文件的请求，需要带上 referer")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			stack := make([]byte, 1024)
			runtime.Stack(stack, true)
			fmt.Println("panic:", r, "\nstack:", string(stack))
			os.Exit(-1)
		}
	}()

	if url == "" {
		panicParameter("u")
	}
	if output == "" {
		panicParameter("o")
	}
	if chanSize <= 0 {
		panic("parameter 'c' must be greater than 0")
	}

	downloader, err := dl.NewTask(output, url, referer)
	if err != nil {
		panic(err)
	}
	if err := downloader.Start(chanSize); err != nil {
		panic(err)
	}
	fmt.Println("Done!")
}

func panicParameter(name string) {
	panic("parameter '" + name + "' is required")
}
