package converter

import (
	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"
)

// schemas.HealthcheckConfigBodyをmodels.HealthcheckConfigに変換
func ConvertHealthcheckConfigBodyToModel(body *schemas.HealthcheckConfigBody) *models.HealthcheckConfig {
	if body == nil {
		return nil
	}
	return &models.HealthcheckConfig{
		SystemID:      body.SystemID,
		Name:          body.Name,
		Description:   body.Description,
		ConfigType:    body.ConfigType,
		ExpectedValue: body.ExpectedValue,
		Url:           body.Url,
		IsActive:      body.IsActive,
	}
}

// models.HealthcheckConfigをschemas.HealthcheckConfigResponseに変換
func ConvertModelToHealthcheckConfigResponse(model *models.HealthcheckConfig) *schemas.HealthcheckConfigResponse {
	if model == nil {
		return nil
	}
	return &schemas.HealthcheckConfigResponse{
		ID:            model.ID,
		SystemID:      model.SystemID,
		Name:          model.Name,
		Description:   model.Description,
		ConfigType:    model.ConfigType,
		ExpectedValue: model.ExpectedValue,
		Url:           model.Url,
		IsActive:      model.IsActive,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// HealthcheckConfigの配列をHealthcheckConfigsResponseに変換
func ConvertHealthcheckConfigsToResponse(system models.System, configs []models.HealthcheckConfig) schemas.HealthcheckConfigsResponse {
	ConfigResponses := []schemas.HealthcheckConfigResponse{}

	for _, config := range configs {
		ConfigResponses = append(ConfigResponses, schemas.HealthcheckConfigResponse{
			ID:            config.ID,
			SystemID:      config.SystemID,
			Name:          config.Name,
			Description:   config.Description,
			ConfigType:    config.ConfigType,
			ExpectedValue: config.ExpectedValue,
			Url:           config.Url,
			IsActive:      config.IsActive,
			CreatedAt:     config.CreatedAt,
			UpdatedAt:     config.UpdatedAt,
		})
	}

	// システムモデルをレスポンススキーマに変換
	systemResponse := ConvertSystemModelToResponseSchema(&system)

	return schemas.HealthcheckConfigsResponse{
		SystemResponse: *systemResponse,
		Configs:        ConfigResponses,
	}
}

// models.HealthcheckConfigをschemas.HealthcheckConfigResponseに変換
func ConvertHealthcheckConfigToResponse(model *models.HealthcheckConfig) *schemas.HealthcheckConfigResponse {
	if model == nil {
		return nil
	}
	return &schemas.HealthcheckConfigResponse{
		ID:            model.ID,
		SystemID:      model.SystemID,
		Name:          model.Name,
		Description:   model.Description,
		ConfigType:    model.ConfigType,
		ExpectedValue: model.ExpectedValue,
		Url:           model.Url,
		IsActive:      model.IsActive,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// models.HealthcheckLogをschemas.HealthcheckLogsResponseに変換
func ConvertModelToHealthcheckLogsResponse(model *models.HealthcheckLog) *schemas.HealthcheckLogsResponse {
	if model == nil {
		return nil
	}
	return &schemas.HealthcheckLogsResponse{
		ID:                  model.ID,
		IsAlive:             model.IsAlive,
		ResponseValue:       model.ResponseValue,
		HealthcheckConfigId: model.HealthcheckConfigId,
		CreatedAt:           model.CreatedAt,
		UpdatedAt:           model.UpdatedAt,
	}
}
