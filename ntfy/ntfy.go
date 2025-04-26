package ntfy

import (
	"bytes"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Publish(server string, topic string, msg Message, auth Auth, client *http.Client) error {
	if server == "" {
		server = "https://ntfy.sh/"
	}
	if r, _ := utf8.DecodeLastRuneInString(server); r != '/' {
		server += "/"
	}

	if msg.Filename != "" && msg.Attach == "" && msg.File == nil {
		return errors.New("no file attached")
	}
	var body *bytes.Reader
	if msg.Attach == "" && msg.File != nil {
		body = bytes.NewReader(msg.File)
	}

	req, err := http.NewRequest(http.MethodPost, server+topic, body)
	if err != nil {
		return err
	}

	if auth != nil {
		req.Header.Add("Authorization", auth.AuthHeader())
	}

	if msg.Title != "" {
		req.Header.Add("X-Title", msg.Title)
	}
	if msg.Message != "" {
		req.Header.Add("X-Message", msg.Message)
	}
	if msg.Priority != PriorityUndefined {
		req.Header.Add("X-Priority", strconv.Itoa(int(msg.Priority)))
	}
	if msg.Tags != nil {
		req.Header.Add("X-Tags", strings.Join(msg.Tags, ", "))
	}
	if msg.Delay != "" {
		req.Header.Add("X-Delay", msg.Delay)
	}
	if msg.Actions != nil {
		actstrs := []string{}
		for _, act := range msg.Actions {
			if act == nil {
				continue
			}
			actstrs = append(actstrs, act.ActionHeader())
		}
		req.Header.Add("X-Actions", strings.Join(actstrs, "; "))
	}
	if msg.Click != "" {
		req.Header.Add("X-Click", msg.Click)
	}
	if msg.Attach != "" {
		req.Header.Add("X-Attach", msg.Attach)
	}
	if msg.Markdown {
		req.Header.Add("X-Markdown", "true")
	}
	if msg.Icon != "" {
		req.Header.Add("X-Icon", msg.Icon)
	}
	if msg.Filename != "" {
		req.Header.Add("X-Filename", msg.Filename)
	}
	if msg.Email != "" {
		req.Header.Add("X-Email", msg.Email)
	}
	if msg.Call != "" {
		req.Header.Add("X-Call", msg.Call)
	}
	if msg.NoCache {
		req.Header.Add("X-Cache", "false")
	}
	if msg.NoFirebase {
		req.Header.Add("X-Firebase", "false")
	}
	if msg.UnifiedPush {
		req.Header.Add("X-UnifiedPush", "true")
	}

	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("publish failed")
	}

	return nil
}
