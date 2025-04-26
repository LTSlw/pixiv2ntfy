package pixiv

import (
	"net/url"
	"time"
)

type Illust struct {
	ID         uint64
	Title      string
	Comment    string
	Pages      []Page
	AuthorID   uint64
	AuthorName string
	CreateDate time.Time
	UploadDate time.Time
}

type Page struct {
	URL    url.URL
	Width  uint64
	Height uint64
}

type pixivIllustResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		IllustID      string `json:"illustId"`
		IllustTitle   string `json:"illustTitle"`
		IllustComment string `json:"illustComment"`
		ID            string `json:"id"`
		Title         string `json:"title"`
		Description   string `json:"description"`
		IllustType    int    `json:"illustType"`
		CreateDate    string `json:"createDate"`
		UploadDate    string `json:"uploadDate"`
		Restrict      int    `json:"restrict"`
		XRestrict     int    `json:"xRestrict"`
		Sl            int    `json:"sl"`
		Urls          struct {
			Mini     string `json:"mini"`
			Thumb    string `json:"thumb"`
			Small    string `json:"small"`
			Regular  string `json:"regular"`
			Original string `json:"original"`
		} `json:"urls"`
		Tags struct {
			AuthorID string `json:"authorId"`
			IsLocked bool   `json:"isLocked"`
			Tags     []struct {
				Tag         string `json:"tag"`
				Locked      bool   `json:"locked"`
				Deletable   bool   `json:"deletable"`
				UserID      string `json:"userId"`
				UserName    string `json:"userName"`
				Translation struct {
					En string `json:"en"`
				} `json:"translation,omitempty"`
			} `json:"tags"`
			Writable bool `json:"writable"`
		} `json:"tags"`
		Alt         string `json:"alt"`
		UserID      string `json:"userId"`
		UserName    string `json:"userName"`
		UserAccount string `json:"userAccount"`
		UserIllusts map[string]struct {
			ID                      string   `json:"id"`
			Title                   string   `json:"title"`
			IllustType              int      `json:"illustType"`
			XRestrict               int      `json:"xRestrict"`
			Restrict                int      `json:"restrict"`
			Sl                      int      `json:"sl"`
			URL                     string   `json:"url"`
			Description             string   `json:"description"`
			Tags                    []string `json:"tags"`
			UserID                  string   `json:"userId"`
			UserName                string   `json:"userName"`
			Width                   int      `json:"width"`
			Height                  int      `json:"height"`
			PageCount               int      `json:"pageCount"`
			IsBookmarkable          bool     `json:"isBookmarkable"`
			BookmarkData            any      `json:"bookmarkData"`
			Alt                     string   `json:"alt"`
			TitleCaptionTranslation struct {
				WorkTitle   any `json:"workTitle"`
				WorkCaption any `json:"workCaption"`
			} `json:"titleCaptionTranslation"`
			CreateDate      string `json:"createDate"`
			UpdateDate      string `json:"updateDate"`
			IsUnlisted      bool   `json:"isUnlisted"`
			IsMasked        bool   `json:"isMasked"`
			AIType          int    `json:"aiType"`
			ProfileImageURL string `json:"profileImageUrl,omitempty"`
		} `json:"userIllusts"`
		LikeData             bool  `json:"likeData"`
		Width                int   `json:"width"`
		Height               int   `json:"height"`
		PageCount            int   `json:"pageCount"`
		BookmarkCount        int   `json:"bookmarkCount"`
		LikeCount            int   `json:"likeCount"`
		CommentCount         int   `json:"commentCount"`
		ResponseCount        int   `json:"responseCount"`
		ViewCount            int   `json:"viewCount"`
		BookStyle            any   `json:"bookStyle"`
		IsHowto              bool  `json:"isHowto"`
		IsOriginal           bool  `json:"isOriginal"`
		ImageResponseOutData []any `json:"imageResponseOutData"`
		ImageResponseData    []any `json:"imageResponseData"`
		ImageResponseCount   int   `json:"imageResponseCount"`
		PollData             any   `json:"pollData"`
		SeriesNavData        any   `json:"seriesNavData"`
		DescriptionBoothID   any   `json:"descriptionBoothId"`
		DescriptionYoutubeID any   `json:"descriptionYoutubeId"`
		ComicPromotion       any   `json:"comicPromotion"`
		FanboxPromotion      any   `json:"fanboxPromotion"`
		ContestBanners       []any `json:"contestBanners"`
		IsBookmarkable       bool  `json:"isBookmarkable"`
		BookmarkData         any   `json:"bookmarkData"`
		ContestData          any   `json:"contestData"`
		ZoneConfig           struct {
			Responsive struct {
				URL string `json:"url"`
			} `json:"responsive"`
			Rectangle struct {
				URL string `json:"url"`
			} `json:"rectangle"`
			Size500x500 struct {
				URL string `json:"url"`
			} `json:"500x500"`
			Header struct {
				URL string `json:"url"`
			} `json:"header"`
			Footer struct {
				URL string `json:"url"`
			} `json:"footer"`
			ExpandedFooter struct {
				URL string `json:"url"`
			} `json:"expandedFooter"`
			Logo struct {
				URL string `json:"url"`
			} `json:"logo"`
			AdLogo struct {
				URL string `json:"url"`
			} `json:"ad_logo"`
			Relatedworks struct {
				URL string `json:"url"`
			} `json:"relatedworks"`
		} `json:"zoneConfig"`
		ExtraData struct {
			Meta struct {
				Title              string `json:"title"`
				Description        string `json:"description"`
				Canonical          string `json:"canonical"`
				AlternateLanguages any    `json:"alternateLanguages"`
				DescriptionHeader  string `json:"descriptionHeader"`
				Ogp                struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Type        string `json:"type"`
				} `json:"ogp"`
				Twitter struct {
					Description string `json:"description"`
					Image       string `json:"image"`
					Title       string `json:"title"`
					Card        string `json:"card"`
				} `json:"twitter"`
			} `json:"meta"`
		} `json:"extraData"`
		TitleCaptionTranslation struct {
			WorkTitle   any `json:"workTitle"`
			WorkCaption any `json:"workCaption"`
		} `json:"titleCaptionTranslation"`
		IsUnlisted           bool `json:"isUnlisted"`
		Request              any  `json:"request"`
		CommentOff           int  `json:"commentOff"`
		AIType               int  `json:"aiType"`
		ReuploadDate         any  `json:"reuploadDate"`
		LocationMask         bool `json:"locationMask"`
		CommissionLinkHidden bool `json:"commissionLinkHidden"`
		IsLoginOnly          bool `json:"isLoginOnly"`
		NoLoginData          struct {
			Breadcrumbs struct {
				Successor []any `json:"successor"`
				Current   struct {
					Zh string `json:"zh"`
				} `json:"current"`
			} `json:"breadcrumbs"`
			ZengoIDWorks []struct {
				ID                      string   `json:"id"`
				Title                   string   `json:"title"`
				IllustType              int      `json:"illustType"`
				XRestrict               int      `json:"xRestrict"`
				Restrict                int      `json:"restrict"`
				Sl                      int      `json:"sl"`
				URL                     string   `json:"url"`
				Description             string   `json:"description"`
				Tags                    []string `json:"tags"`
				UserID                  string   `json:"userId"`
				UserName                string   `json:"userName"`
				Width                   int      `json:"width"`
				Height                  int      `json:"height"`
				PageCount               int      `json:"pageCount"`
				IsBookmarkable          bool     `json:"isBookmarkable"`
				BookmarkData            any      `json:"bookmarkData"`
				Alt                     string   `json:"alt"`
				TitleCaptionTranslation struct {
					WorkTitle   any `json:"workTitle"`
					WorkCaption any `json:"workCaption"`
				} `json:"titleCaptionTranslation"`
				CreateDate      string `json:"createDate"`
				UpdateDate      string `json:"updateDate"`
				IsUnlisted      bool   `json:"isUnlisted"`
				IsMasked        bool   `json:"isMasked"`
				AIType          int    `json:"aiType"`
				ProfileImageURL string `json:"profileImageUrl,omitempty"`
			} `json:"zengoIdWorks"`
			ZengoWorkData struct {
				NextWork struct {
					ID    string `json:"id"`
					Title string `json:"title"`
				} `json:"nextWork"`
				PrevWork struct {
					ID    string `json:"id"`
					Title string `json:"title"`
				} `json:"prevWork"`
			} `json:"zengoWorkData"`
		} `json:"noLoginData"`
	} `json:"body"`
}

type pixivIllustPagesResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    []struct {
		Urls struct {
			ThumbMini string `json:"thumb_mini"`
			Small     string `json:"small"`
			Regular   string `json:"regular"`
			Original  string `json:"original"`
		} `json:"urls"`
		Width  uint64 `json:"width"`
		Height uint64 `json:"height"`
	} `json:"body"`
}
