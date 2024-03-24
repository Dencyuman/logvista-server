package background

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// 指定したurlから返ってくるhtmlのtitle属性を取得する関数
func FetchPageTitle(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	title := doc.Find("title").Text()
	return title, nil
}

// 指定したurl(エンドポイント)から取得されるStatusCodeを取得する関数
func FetchStatusCode(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
