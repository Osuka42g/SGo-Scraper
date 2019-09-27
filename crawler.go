package main

import (
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func crawlImages(rawContents io.Reader) []string {
	z := html.NewTokenizer(rawContents)
	imagesFound := make([]string, 0)

	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return imagesFound
		case tt == html.StartTagToken:
			t := z.Token()
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}
			link := getValueFromAttribute(t, "href")
			if link == "" {
				continue
			}
			hasProto := strings.Index(link, "https://") == 0 && strings.HasSuffix(link, ".jpg") == true
			if hasProto {
				imagesFound = append(imagesFound, link)
			}
		}
	}
}

func getAlbumInfo(rawContents io.Reader) (modelName string, albumName string) {
	title := getTitle(rawContents)
	s := strings.Split(title, " Photo Album: ")
	ss := strings.Split(s[1], " | SuicideGirls")
	return strings.TrimSpace(s[0]), strings.TrimSpace(ss[0])
}

func getTitle(rawContents io.Reader) string {
	z := html.NewTokenizer(rawContents)
	defaultTitle := ""
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return defaultTitle
		case tt == html.StartTagToken:
			t := z.Token()
			isTitle := t.Data == "title"
			if !isTitle {
				continue
			}
			z.Next()
			title := z.Token()
			return title.Data
		}
	}
}

func getContents(link string) io.Reader {
	sessionId := settings[settingsSessionId]

	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "sessid",
		Value:  sessionId,
		Path:   "/",
		Domain: "www.suicidegirls.com",
	}

	cookies = append(cookies, cookie)

	u, _ := url.Parse(link)
	jar.SetCookies(u, cookies)

	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", link, nil)
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return resp.Body
}

func getValueFromAttribute(t html.Token, attr string) string {
	val := ""
	for _, a := range t.Attr {
		if a.Key == attr {
			val = a.Val
		}
	}

	return val
}
