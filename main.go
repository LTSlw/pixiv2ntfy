package main

import (
	"os"

	"github.com/LTSlw/pixiv2ntfy/pixiv"
)

func main() {
	pics, err := pixiv.Download(129138049)
	_ = err
	f, _ := os.Create("/tmp/res")
	f.Write(pics[0])
}
