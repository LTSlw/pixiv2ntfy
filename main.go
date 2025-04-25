package main

import (
	"fmt"

	"github.com/LTSlw/pixiv2ntfy/ntfy"
	"github.com/LTSlw/pixiv2ntfy/pixiv"
)

func main() {
	pics, err := pixiv.Download(126150365)
	if err != nil {
		fmt.Println(err.Error())
	}
	msg := &ntfy.Message{
		Title: "Pictures Today",
		File:  pics[0],
	}
	ntfy.Publish("", "pixiv2ntfy-test", *msg, nil, nil)
}
