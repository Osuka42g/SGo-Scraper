package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func crawlImages(rawContents io.ReadCloser) []string {

	z := html.NewTokenizer(rawContents)
	imagesFound := []string{}

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

// todo
func getAlbumName(rawContents io.ReadCloser) string {
	z := html.NewTokenizer(rawContents)
	albumName := ""
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return albumName
		case tt == html.StartTagToken:
			t := z.Token()
			isHeader := t.Data == "h2"
			if !isHeader {
				continue
			}
			if getValueFromAttribute(t, "class") == "title" {

			}
			fmt.Println(t)
		}
	}
}

func getModelName(rawContents io.ReadCloser) string {
	return "model"
}

func getContents(link string) io.ReadCloser {
	sessionidCookie := os.Getenv("SESSIONIDTOKEN")

	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie
	cookie := &http.Cookie{
		Name:   "sessionid",
		Value:  sessionidCookie,
		Path:   "/",
		Domain: "www.suicidegirls.com",
	}

	cookies = append(cookies, cookie)

	u, _ := url.Parse(link)
	jar.SetCookies(u, cookies)
	fmt.Println(jar.Cookies(u))

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
