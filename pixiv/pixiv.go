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

func Download(pid uint64) ([][]byte, error) {
	illust, err := GetIllust(pid)
	if err != nil {
		return nil, err
	}

	return illust.Download()
}

func (i *Illust) Download() ([][]byte, error) {
	pics := [][]byte{}
	for _, p := range i.Pages {
		pic, err := download(http.DefaultClient, p.URL.String(), "", imageDownloadReferer)
		if err != nil {
			return nil, err
		}
		pics = append(pics, pic)
	}
	return pics, nil
}

func GetIllust(pid uint64) (*Illust, error) {
	info, err := getIllustInfo(pid)
	if err != nil {
		return nil, err
	}

	pages, err := getIllustPages(pid)
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

func getIllustInfo(pid uint64) (*pixivIllustResponse, error) {
	raw, err := download(http.DefaultClient, fmt.Sprintf(illustURL, pid), "", "")
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

func getIllustPages(pid uint64) (*pixivIllustPagesResponse, error) {
	raw, err := download(http.DefaultClient, fmt.Sprintf(illustPagesURL, pid), "", "")
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

func download(client *http.Client, url, cookie, referer string) ([]byte, error) {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	if referer != "" {
		req.Header.Add("Referer", referer)
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

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
