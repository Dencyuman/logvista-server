package background

import (
	"fmt"
	"github.com/Dencyuman/logvista-server/config"
	"github.com/Dencyuman/logvista-server/src/crud"
	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/gorm"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type BackgroundApp struct {
	DB *gorm.DB
}

// タイムアウトを含むHTTPクライアントの設定
var httpClient = &http.Client{
	Timeout: 10 * time.Second, // 例として10秒のタイムアウトを設定
}

// 指定したurlから返ってくるhtmlのtitle属性を取得する関数
func FetchPageTitle(url string) (string, error) {
	res, err := httpClient.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

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

// 指定したURLのエンドポイントから返ってくるレスポンスボディを文字列で取得する関数
func FetchHealthcheckAPIResponseAsString(url string) (string, error) {
	res, err := httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// ログ出力などエラー処理をここに追加
		}
	}(res.Body)

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(bodyBytes), nil
}

// 60秒ごとにサーバーのヘルスチェックを行う関数
func (ctrl *BackgroundApp) checkServer() {
	interval := time.Duration(config.AppConfig.HealthcheckTimespan) * time.Second
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Print("Checking server health...")
		configurations, err := crud.FindAllActiveHealthcheckConfigs(ctrl.DB)
		if err != nil {
			log.Printf("設定の取得中にエラーが発生しました: %v", err)
			continue
		}

		var waitGroup sync.WaitGroup
		for _, configuration := range configurations {
			waitGroup.Add(1)
			go func(config models.HealthcheckConfig) {
				defer waitGroup.Done()
				if config.ConfigType == models.SiteTitle {
					title, titleError := FetchPageTitle(config.Url)
					if titleError != nil {
						log.Printf("URL %s のタイトル取得中にエラー: %v", config.Url, titleError)
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             false,
							ResponseValue:       titleError.Error(),
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					}
					if title != config.ExpectedValue {
						log.Printf("URL %s の期待されるタイトル '%s' と異なります。取得されたタイトル: '%s'", config.Url, config.ExpectedValue, title)
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             false,
							ResponseValue:       title,
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					} else {
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             true,
							ResponseValue:       title,
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					}
				} else if config.ConfigType == models.Endpoint {
					response, FetchError := FetchHealthcheckAPIResponseAsString(config.Url)
					if FetchError != nil {
						log.Printf("URL %s のステータスコード取得中にエラー: %v", config.Url, FetchError)
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             false,
							ResponseValue:       FetchError.Error(),
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					}
					if fmt.Sprintf("%s", response) != config.ExpectedValue {
						log.Printf("URL %s の期待されるステータスコード '%s' と異なります。取得されたコード: '%s'", config.Url, config.ExpectedValue, response)
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             false,
							ResponseValue:       fmt.Sprintf("%s", response),
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					} else {
						healthcheckLog := &models.HealthcheckLog{
							IsAlive:             true,
							ResponseValue:       fmt.Sprintf("%s", response),
							HealthcheckConfigId: config.ID,
						}
						if err := crud.InsertHealthcheckLog(ctrl.DB, healthcheckLog); err != nil {
							log.Printf("HealthcheckLogの挿入中にエラー: %v", err)
						}
						return
					}
				}
			}(configuration)
		}
		waitGroup.Wait()
	}
}

func SetupServerChecker(db *gorm.DB) {
	app := BackgroundApp{
		DB: db,
	}
	fmt.Print("Starting server checker...")
	go app.checkServer()
}
