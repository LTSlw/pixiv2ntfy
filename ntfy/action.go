package ntfy

import (
	"fmt"
	"strings"
	"unicode"
)

type ActionView struct {
	Label string
	URL   string
	Clear bool
}

func (a *ActionView) ActionHeader() string {
	return fmt.Sprintf("view, %s, %s, clear=%t", escape(a.Label), escape(a.URL), a.Clear)
}

type ActionBroadcast struct {
	Label  string
	Intent string
	Extras map[string]string
	Clear  bool
}

func (a *ActionBroadcast) ActionHeader() string {
	actstr := fmt.Sprintf("broadcast, %s, clear=%t", escape(a.Label), a.Clear)
	if a.Intent != "" {
		actstr += fmt.Sprintf(", intent=%s", escape(a.Intent))
	}
	for k, v := range a.Extras {
		actstr += fmt.Sprintf(", extras.%s=%s", toKeyStr(k), escape(v))
	}
	return actstr
}

type ActionHTTP struct {
	Label   string
	URL     string
	Method  string
	Headers map[string]string
	Body    string
	Clear   bool
}

func (a *ActionHTTP) ActionHeader() string {
	actstr := fmt.Sprintf("http, %s, %s, clear=%t", escape(a.Label), escape(a.URL), a.Clear)
	if a.Method != "" {
		actstr += fmt.Sprintf(", method=%s", escape(a.Method))
	}
	for k, v := range a.Headers {
		actstr += fmt.Sprintf(", headers.%s=%s", toKeyStr(k), escape(v))
	}
	if a.Body != "" {
		actstr += fmt.Sprintf(", body=%s", escape(a.Body))
	}

	return actstr
}

type Action interface {
	ActionHeader() string
}

func toKeyStr(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) || unicode.IsLetter(r) || r == '_' || r == '-' {
			return r
		}
		return -1
	}, s)
}

func escape(s string) string {
	if !strings.ContainsRune(s, '\'') && !strings.ContainsRune(s, '=') {
		return s
	}
	s = strings.ReplaceAll(s, "'", "\\'")
	return fmt.Sprintf("'%s'", s)
}
