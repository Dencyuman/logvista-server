package schemas

import "time"

// swagger:request testHealthcheckSiteTitleConfigBody
type TestHealthcheckSiteTitleConfigBody struct {
	ExpectedTitle string `json:"expected_title" example:"sampleTitle"` // 想定タイトル
	Url           string `json:"url" example:"http://localhost:8080/"` // アクセス先url
}

// swagger:response testHealthcheckSiteTitleConfigResponse
type TestHealthcheckSiteTitleConfigResponse struct {
	ExpectedTitle     string `json:"expected_title" example:"sampleTitle"` // 想定タイトル
	FetchedTitle      string `json:"fetched-title" example:"sampleTitle"`  // 取得されたタイトル
	HealthcheckResult bool   `json:"healthcheck_result" example:"true"`    // ヘルスチェック結果
}

// swagger:request testHealthcheckEndpointConfigBody
type TestHealthcheckEndpointConfigBody struct {
	ExpectedStatus string `json:"expected_status" example:"sampleSystem API Server is running!"` // 想定レスポンス
	Url            string `json:"url" example:"http://localhost:8080/"`                          // アクセス先url
}

// swagger:response testHealthcheckEndpointConfigResponse
type TestHealthcheckEndpointConfigResponse struct {
	ExpectedStatus    string `json:"expected_status" example:"sampleSystem API Server is running!"` // 想定レスポンス
	FetchedStatus     string `json:"fetched-status" example:"sampleStatus"`                         // 取得されたタイトル
	HealthcheckResult bool   `json:"healthcheck_result" example:"true"`                             // ヘルスチェック結果
}

// swagger:request healthcheckSiteTitleConfigBody
type HealthcheckSiteTitleConfigBody struct {
	SystemID      string `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name          string `json:"name" example:"sampleName"`                                                   // 設定名
	Description   string `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedTitle string `json:"expected_title" example:"sampleTitle"`                                        // 想定タイトル
	Url           string `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	IsActive      bool   `json:"is_active" example:"true"`                                                    // 有効かどうか
}

// swagger:response healthcheckSiteTitleConfigResponse
type HealthcheckSiteTitleConfigResponse struct {
	ID            string    `json:"id" example:"00000000-0000-0000-0000-000000000000"`                           // ログID
	SystemID      string    `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name          string    `json:"name" example:"sampleName"`                                                   // 設定名
	Description   string    `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedTitle string    `json:"expected_title" example:"sampleTitle"`                                        // 想定タイトル
	Url           string    `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	IsActive      bool      `json:"is_active" example:"true"`                                                    // 有効かどうか
	CreatedAt     time.Time `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt     time.Time `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:request healthcheckSiteTitleConfigBody
type HealthcheckEndpointConfigBody struct {
	SystemID         string `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name             string `json:"name" example:"sampleName"`                                                   // 設定名
	Description      string `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedResponse string `json:"expected_response" example:"sampleSystem API Server is running!"`             // 想定レスポンス
	Url              string `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	IsActive         bool   `json:"is_active" example:"true"`                                                    // 有効かどうか
}

// swagger:response healthcheckSiteTitleConfigResponse
type HealthcheckEndpointConfigResponse struct {
	ID               string    `json:"id" example:"00000000-0000-0000-0000-000000000000"`                           // ログID
	SystemID         string    `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name             string    `json:"name" example:"sampleName"`                                                   // 設定名
	Description      string    `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ExpectedResponse string    `json:"expected_response" example:"sampleSystem API Server is running!"`             // 想定レスポンス
	Url              string    `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	IsActive         bool      `json:"is_active" example:"true"`                                                    // 有効かどうか
	CreatedAt        time.Time `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt        time.Time `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:response healthcheckConfigsResponse
type HealthcheckConfigsResponse struct {
	SystemResponse                                        // システム情報
	SiteTitleConfigs []HealthcheckSiteTitleConfigResponse `json:"site_title_configs"` // サイトタイトルヘルスチェックの設定
	EndpointConfigs  []HealthcheckEndpointConfigResponse  `json:"endpoint_configs"`   // エンドポイントヘルスチェックの設定
}

// swagger:response healthcheckConfigResponse
type HealthcheckConfigResponse struct {
	ID            string    `json:"id" example:"00000000-0000-0000-0000-000000000000"`        // 設定ID
	SystemID      string    `json:"system_id" example:"00000000-0000-0000-0000-000000000000"` // システムID
	Name          string    `json:"name" example:"sampleName"`                                // 設定名
	Description   string    `json:"description" example:"sampleDescription"`                  // 設定の説明
	ConfigType    string    `json:"config_type" example:"SiteTitle"`                          // 設定タイプ
	ExpectedValue string    `json:"expected_value" example:"sampleValue"`                     // 想定値
	Url           string    `json:"url" example:"http://localhost:8080/"`                     // アクセス先url
	IsActive      bool      `json:"is_active" example:"true"`                                 // 有効かどうか
	CreatedAt     time.Time `json:"created_at" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt     time.Time `json:"updated_at" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:response healthcheckLogsResponse
type HealthcheckLogsResponse struct {
	ID                  string    `json:"id" example:"00000000-0000-0000-0000-000000000000"`                    // ログID
	IsAlive             bool      `json:"is_alive" example:"true"`                                              // ヘルスチェック結果
	ResponseValue       string    `json:"response_value" example:"sampleResponse"`                              // レスポンス値
	HealthcheckConfigId string    `json:"healthcheck_config_id" example:"00000000-0000-0000-0000-000000000000"` // ヘルスチェック設定ID
	CreatedAt           time.Time `json:"created_at" example:"2023-01-01T00:00:00.000000+09:00"`                // 作成日時
	UpdatedAt           time.Time `json:"updated_at" example:"2023-01-01T00:00:00.000000+09:00"`                // 更新日時
}

// swagger:response healthcheckLogsListResponse
type HealthcheckLogsListResponse struct {
	Config HealthcheckConfigResponse `json:"config"` // ヘルスチェック設定
	Logs   []HealthcheckLogsResponse `json:"logs"`   // ログリスト
}
