package pixiv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const illustURL = "https://www.pixiv.net/ajax/illust/%d"
const illustPagesURL = "https://www.pixiv.net/ajax/illust/%d/pages"
const imageDownloadReferer = "https://pixiv.net"
const defaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:137.0) Gecko/20100101 Firefox/114514.0"

func Download(pid uint64, sessid, ua string) ([][]byte, error) {
	illust, err := GetIllust(pid, sessid, ua)
	if err != nil {
		return nil, err
	}

	return illust.Download()
}

func (i *Illust) Download() ([][]byte, error) {
	pics := [][]byte{}
	for _, p := range i.Pages {
		pic, err := download(http.DefaultClient, p.URL.String(), imageDownloadReferer, "", "")
		if err != nil {
			return nil, err
		}
		pics = append(pics, pic)
	}
	return pics, nil
}

func (i *Illust) DownloadPage(pageID uint) ([]byte, error) {
	return download(http.DefaultClient, i.Pages[pageID].URL.String(), imageDownloadReferer, "", "")
}

func GetIllust(pid uint64, sessid, ua string) (*Illust, error) {
	info, err := getIllustInfo(pid, sessid, ua)
	if err != nil {
		return nil, err
	}

	pages, err := getIllustPages(pid, sessid, ua)
	if err != nil {
		return nil, err
	}

	illust := &Illust{
		ID:      unwarp(strconv.ParseUint(info.Body.IllustID, 10, 64)),
		Title:   info.Body.IllustTitle,
		Comment: info.Body.IllustComment,
		Pages: func() []Page {
			ps := []Page{}
			for _, p := range pages.Body {
				page := Page{
					URL:    *unwarp(url.Parse(p.Urls.Original)),
					Width:  p.Width,
					Height: p.Height,
				}
				ps = append(ps, page)
			}
			return ps
		}(),
		AuthorID:   unwarp(strconv.ParseUint(info.Body.UserID, 10, 64)),
		AuthorName: info.Body.UserName,
		CreateDate: unwarp(time.Parse(time.RFC3339, info.Body.CreateDate)),
		UploadDate: unwarp(time.Parse(time.RFC3339, info.Body.UploadDate)),
	}
	return illust, nil
}

func getIllustInfo(pid uint64, sessid, ua string) (*pixivIllustResponse, error) {
	raw, err := download(http.DefaultClient, fmt.Sprintf(illustURL, pid), "", sessid, ua)
	if err != nil {
		return nil, err
	}

	resp, err := parseJSON[pixivIllustResponse](raw)
	if err != nil {
		return nil, err
	}

	if resp.Error {
		return nil, errors.New("get illust info failed")
	}

	return resp, nil
}

func getIllustPages(pid uint64, sessid, ua string) (*pixivIllustPagesResponse, error) {
	raw, err := download(http.DefaultClient, fmt.Sprintf(illustPagesURL, pid), "", sessid, ua)
	if err != nil {
		return nil, err
	}

	resp, err := parseJSON[pixivIllustPagesResponse](raw)
	if err != nil {
		return nil, err
	}

	if resp.Error {
		return nil, errors.New("get illust pages failed")
	}
	return resp, nil
}

func download(client *http.Client, url, referer, sessid, ua string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if sessid != "" {
		c := &http.Cookie{
			Name:  "PHPSESSID",
			Value: sessid,
		}
		req.AddCookie(c)
		if ua == "" {
			req.Header.Add("User-Agent", defaultUserAgent)
		}
	}
	if ua != "" {
		req.Header.Add("User-Agent", ua)
	}
	if referer != "" {
		req.Header.Add("Referer", referer)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http response is not status ok")
	}

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func parseJSON[T any](raw []byte) (*T, error) {
	t := new(T)
	err := json.Unmarshal(raw, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func unwarp[T any](t T, _ error) T {
	return t
}

func ptr[T any](t T) *T {
	return &t
}
