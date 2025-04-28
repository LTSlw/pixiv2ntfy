package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/LTSlw/pixiv2ntfy/ntfy"
	"github.com/LTSlw/pixiv2ntfy/pixiv"
)

var (
	pixivSessID  string
	pixivUA      string
	pid          uint64
	pageID       uint64
	ntfyServer   string
	ntfyTopic    string
	ntfyToken    string
	ntfyUsername string
	ntfyPassword string
	ntfyAuth     ntfy.Auth
)

func init() {
	bindStr(&pixivSessID, "P2N_PIXIV_SESSION_ID", "S", "", "Pixiv session id")
	bindStr(&pixivUA, "P2N_PIXIV_USER_AGENT", "UA", "", "Pixiv user agent")
	bindUint64(&pageID, "", "p", 0, "Page number of illust")
	bindStr(&ntfyServer, "P2N_NTFY_SERVER", "N", "", "Ntfy server base url")
	bindStr(&ntfyTopic, "P2N_NTFY_TOPIC", "topic", "", "Ntfy topic")
	bindStr(&ntfyToken, "P2N_NTFY_TOKEN", "T", "", "Ntfy token")
	bindStr(&ntfyUsername, "P2N_NTFY_USERNAME", "U", "", "Ntfy username")
	bindStr(&ntfyPassword, "P2N_NTFY_PASSWORD", "P", "", "Ntfy password")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: pixiv2ntfy [FLAGS...] [TOPIC] PIXIV_ID\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	pidstr := ""
	switch flag.NArg() {
	case 0:
		failf(os.Stderr, "no enough arguments\n")
	case 1:
		if ntfyTopic == "" {
			failf(os.Stderr, "ntfy topic is required")
		}
		pidstr = flag.Arg(0)
	case 2:
		ntfyTopic = flag.Arg(0)
		pidstr = flag.Arg(1)
	default:
		failf(os.Stderr, "too many arguments\n")
	}
	var err error
	pid, err = strconv.ParseUint(pidstr, 10, 64)
	if err != nil {
		failf(os.Stderr, "invalid value %q for argument pid: parse error\n", pidstr)
	}

	if ntfyUsername != "" {
		ntfyAuth = &ntfy.AuthUserPassword{
			Username: ntfyUsername,
			Password: ntfyPassword,
		}
	}
	if ntfyToken != "" {
		ntfyAuth = &ntfy.AuthToken{
			Token: ntfyToken,
		}
	}
}

func main() {
	illust, err := pixiv.GetIllust(pid, pixivSessID, pixivUA)
	if err != nil {
		log.Fatalf("get illust details failed: %s\n", err.Error())
	}
	log.Printf("pid: %d\ntitle: %s\ncomment: %s\nauthor: %s\nnum of pages: %d\n", illust.ID, illust.Title, illust.Comment, illust.AuthorName, len(illust.Pages))

	img, err := illust.DownloadPage(uint(pageID))
	if err != nil {
		log.Fatalf("download image failed\n")
	}
	log.Println("image downloaded")

	msg := &ntfy.Message{
		Title: "Pixiv Artwork",
		File:  img,
	}

	err = ntfy.Publish(ntfyServer, ntfyTopic, *msg, ntfyAuth, nil)
	if err != nil {
		log.Fatalf("publish message failed: %s\n", err.Error())
	}
	log.Println("message published")
}
