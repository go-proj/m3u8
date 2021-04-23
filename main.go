package main

import (
	"flag"
	"fmt"
	"os"

	"coding.net/cy9/golang/m3u8/dl"
)

var (
	url      string
	output   string
	chanSize int
	tsPort   int
)

// 实际遇到过 key 和 ts 文件，端口可能跟 m3u8 文件不同
func init() {
	flag.StringVar(&url, "u", "", "M3U8 URL, required")
	flag.StringVar(&output, "o", "", "Output folder, required")
	flag.IntVar(&chanSize, "c", 32, "Maximum number of occurrences")
	flag.IntVar(&tsPort, "t", 0, "实际遇到过 key 和 ts 文件，端口可能跟 m3u8 文件不同")
}

func main() {
	flag.Parse()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[error]", r)
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

	downloader, err := dl.NewTask(output, url)
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
