package aocclient

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"
)

var jar http.CookieJar

func init() {
	var err error
	jar, err = cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	u, err := url.Parse("https://adventofcode.com")
	if err != nil {
		panic(err)
	}

	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "session",
			Value: aocCookie,
		},
	})
}

func httpClient() *http.Client {
	return &http.Client{
		Jar: jar,
	}
}

func SubmitAnswer(year, day, part int, answer string) bool {
	uri := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	resp, err := httpClient().PostForm(uri, url.Values{"level": {fmt.Sprint(part)}, "answer": {answer}})
	if err != nil {
		panic(err)
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf))
	return strings.Contains(string(buf), "That's the right answer")
}

func GetInput(year, day int) io.Reader {
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	resp, err := httpClient().Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic("failed to get input")
	}
	return resp.Body
}
