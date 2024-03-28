package schemas

import (
	"github.com/Dencyuman/logvista-server/src/models"
	"time"
)

// swagger:request testHealthcheckConfigBody
type TestHealthcheckConfigBody struct {
	ConfigType    models.HealthcheckConfigType `json:"config_type" binding:"required" example:"SiteTitle"` // 設定タイプ
	ExpectedValue string                       `json:"expected_value" example:"sampleValue"`               // 想定値
	Url           string                       `json:"url" example:"http://localhost:8080/"`               // アクセス先url
}

// swagger:response testHealthcheckConfigResponse
type TestHealthcheckConfigResponse struct {
	ConfigType        models.HealthcheckConfigType `json:"config_type" binding:"required" example:"SiteTitle"` // 設定タイプ
	ExpectedValue     string                       `json:"expected_value" example:"sampleValue"`               // 想定値
	FetchedValue      string                       `json:"fetched-value" example:"sampleValue"`                // 取得された値
	HealthcheckResult bool                         `json:"healthcheck_result" example:"true"`                  // ヘルスチェック結果
}

// swagger:request healthcheckConfigBody
type HealthcheckConfigBody struct {
	SystemID      string                       `json:"system_id" binding:"required" example:"00000000-0000-0000-0000-000000000000"` // システムid
	Name          string                       `json:"name" example:"sampleName"`                                                   // 設定名
	Description   string                       `json:"description" example:"sampleDescription"`                                     // 設定の説明
	ConfigType    models.HealthcheckConfigType `json:"config_type" binding:"required" example:"SiteTitle"`                          // 設定タイプ
	ExpectedValue string                       `json:"expected_value" example:"sampleValue"`                                        // 想定値
	Url           string                       `json:"url" example:"http://localhost:8080/"`                                        // アクセス先url
	IsActive      bool                         `json:"is_active" example:"true"`                                                    // 有効かどうか
}

// swagger:response healthcheckConfigResponse
type HealthcheckConfigResponse struct {
	ID            string                       `json:"id" example:"00000000-0000-0000-0000-000000000000"`        // 設定ID
	SystemID      string                       `json:"system_id" example:"00000000-0000-0000-0000-000000000000"` // システムID
	Name          string                       `json:"name" example:"sampleName"`                                // 設定名
	Description   string                       `json:"description" example:"sampleDescription"`                  // 設定の説明
	ConfigType    models.HealthcheckConfigType `json:"config_type" binding:"required" example:"SiteTitle"`       // 設定タイプ
	ExpectedValue string                       `json:"expected_value" example:"sampleValue"`                     // 想定値
	Url           string                       `json:"url" example:"http://localhost:8080/"`                     // アクセス先url
	IsActive      bool                         `json:"is_active" example:"true"`                                 // 有効かどうか
	CreatedAt     time.Time                    `json:"created_at" example:"2023-01-01T00:00:00.000000+09:00"`    // 作成日時
	UpdatedAt     time.Time                    `json:"updated_at" example:"2023-01-01T00:00:00.000000+09:00"`    // 更新日時
}

// swagger:response healthcheckConfigsResponse
type HealthcheckConfigsResponse struct {
	SystemResponse                             // システム情報
	Configs        []HealthcheckConfigResponse `json:"configs"` // Config
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
