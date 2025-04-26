package pixiv_test

import (
	"net/http"
	"testing"

	"github.com/LTSlw/pixiv2ntfy/pixiv"
)

func TestDownload(t *testing.T) {
	_, err := pixiv.Download_(http.DefaultClient, "https://pixiv.net", "", "", "")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestGetIllustPages(t *testing.T) {
	pages, err := pixiv.GetIllustPages(129138049, "", "")
	if err != nil {
		t.Error(err.Error())
	}
	if pages.Body[0].Urls.Original != "https://i.pximg.net/img-original/img/2025/04/10/00/00/14/129138049_p0.png" {
		t.Errorf("incorrect data: %v", pages)
	}
}

func TestGetIllustInfo(t *testing.T) {
	info, err := pixiv.GetIllustInfo(129138049, "", "")
	if err != nil {
		t.Error(err.Error())
	}
	if info.Body.UserID != "26429571" {
		t.Errorf("incorrect data: %v", info)
	}
}

func TestGetIllust(t *testing.T) {
	illust, err := pixiv.GetIllust(129138049, "", "")
	if err != nil {
		t.Error(err.Error())
	}
	if illust.AuthorID != 26429571 {
		t.Errorf("incorrect data: %v", illust)
	}

	imgs, err := illust.Download()
	if err != nil {
		t.Error(err.Error())
	}
	if imgs[0][0] != 0x89 { // png magic number
		t.Error("not a valid png")
	}
}
