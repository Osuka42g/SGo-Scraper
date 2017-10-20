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

func crawlImages(link string) []string {
	rawContents := getContents(link)

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
			link := getHref(t)
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
func getAlbumName() string {
	return "album"
}

func getModelName() string {
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

func getHref(t html.Token) string {
	href := ""
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
		}
	}

	return href
}
