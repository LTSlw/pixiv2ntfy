package main

import (
	"fmt"

	"github.com/LTSlw/pixiv2ntfy/ntfy"
	"github.com/LTSlw/pixiv2ntfy/pixiv"
)

func main() {
	pics, err := pixiv.Download(127826168, "", "")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	msg := &ntfy.Message{
		Title: "Pictures Today",
		File:  pics[0],
	}
	err = ntfy.Publish("", "pixiv2ntfy-test", *msg, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
