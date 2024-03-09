package schemas

import "time"

// swagger:request testHealthcheckSiteTitleConfigBody
type TestHealthcheckSiteTitleConfigBody struct {
	ExpectedTitle string `json:"expexted_title" example:"sampleTitle"` // 想定タイトル
	Url           string `json:"url" example:"http://localhost:8080/"` // アクセス先url
}

// swagger:response testHealthcheckSiteTitleConfigResponse
type TestHealthcheckSiteTitleConfigResponse struct {
	ExpectedTitle     string `json:"expexted_title" example:"sampleTitle"` // 想定タイトル
	FetchedTitle      string `json:"fetched-title" example:"sampleTitle"`  // 取得されたタイトル
	HealthcheckResult bool   `json:"healthcheck_result" example:"true"`    // ヘルスチェック結果
}

// swagger:request testHealthcheckEndpointConfigBody
type TestHealthcheckEndpointConfigBody struct {
	ExpectedStatus int    `json:"expexted_status" example:"200"`        // 想定ステータス
	Url            string `json:"url" example:"http://localhost:8080/"` // アクセス先url
}

// swagger:response testHealthcheckEndpointConfigResponse
type TestHealthcheckEndpointConfigResponse struct {
	ExpectedStatus    int    `json:"expexted_status" example:"200"`         // 想定ステータス
	FetchedStatus     string `json:"fetched-status" example:"sampleStatus"` // 取得されたタイトル
	HealthcheckResult bool   `json:"healthcheck_result" example:"true"`     // ヘルスチェック結果
}

// swagger:request healthcheckSiteTitleConfigBody
type HealthcheckSiteTitleConfigBody struct {
	SystemID      string `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name          string `json:"name" example:"sampleName"`                                                   // 設定名
	Description   string `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedTitle string `json:"expexted_title" example:"sampleTitle"`                                        // 想定タイトル
	Url           string `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	Timespan      int    `json:"timespan" example:"10"`                                                       // ヘルスチェック間隔(10s)
}

// swagger:response healthcheckSiteTitleConfigResponse
type HealthcheckSiteTitleConfigResponse struct {
	SystemID      string    `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name          string    `json:"name" example:"sampleName"`                                                   // 設定名
	Description   string    `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedTitle string    `json:"expexted_title" example:"sampleTitle"`                                        // 想定タイトル
	Url           string    `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	Timespan      int       `json:"timespan" example:"10"`                                                       // ヘルスチェック間隔(10s)
	CreatedAt     time.Time `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt     time.Time `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:request healthcheckSiteTitleConfigBody
type HealthcheckEndpointConfigBody struct {
	SystemID       string `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name           string `json:"name" example:"sampleName"`                                                   // 設定名
	Description    string `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedStatus int    `json:"expexted_status" example:"200"`                                               // 想定ステータス
	Url            string `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	Timespan       int    `json:"timespan" example:"10"`                                                       // ヘルスチェック間隔(10s)
}

// swagger:response healthcheckSiteTitleConfigResponse
type HealthcheckEndpointConfigResponse struct {
	SystemID       string    `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name           string    `json:"name" example:"sampleName"`                                                   // 設定名
	Description    string    `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedStatus int       `json:"expexted_status" example:"200"`                                               // 想定ステータス
	Url            string    `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	Timespan       int       `json:"timespan" example:"10"`                                                       // ヘルスチェック間隔(10s)
	CreatedAt      time.Time `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt      time.Time `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:response healthcheckConfigsResponse
type HealthcheckConfigsResponse struct {
	SystemResponse                                        // システム情報
	SiteTitleConfigs []HealthcheckSiteTitleConfigResponse `json:"site_title_configs"` // サイトタイトルヘルスチェックの設定
	EndpointConfigs  []HealthcheckEndpointConfigResponse  `json:"endpoint_configs"`   // エンドポイントヘルスチェックの設定
}
