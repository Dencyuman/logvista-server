package converter

import (
	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"
)

// schemas.HealthcheckSiteTitleConfigBodyをmodels.HealthcheckConfigに変換
func ConvertHealthcheckSiteTitleConfigBodyToModel(body *schemas.HealthcheckSiteTitleConfigBody) *models.HealthcheckConfig {
	if body == nil {
		return nil
	}
	return &models.HealthcheckConfig{
		SystemID:      body.SystemID,
		Name:          body.Name,
		Description:   body.Description,
		ConfigType:    models.SiteTitle,
		ExpectedValue: body.ExpectedTitle,
		Url:           body.Url,
		IsActive:      body.IsActive,
	}
}

// schemas.HealthcheckEndpointConfigBodyをmodels.HealthcheckConfigに変換
func ConvertHealthcheckEndpointConfigBodyToModel(body *schemas.HealthcheckEndpointConfigBody) *models.HealthcheckConfig {
	if body == nil {
		return nil
	}
	return &models.HealthcheckConfig{
		SystemID:      body.SystemID,
		Name:          body.Name,
		Description:   body.Description,
		ConfigType:    models.Endpoint,
		ExpectedValue: body.ExpectedResponse,
		Url:           body.Url,
		IsActive:      body.IsActive,
	}
}

// models.HealthcheckConfigをschemas.HealthcheckSiteTitleConfigResponseに変換
func ConvertModelToHealthcheckSiteTitleConfigResponse(model *models.HealthcheckConfig) *schemas.HealthcheckSiteTitleConfigResponse {
	if model == nil {
		return nil
	}
	return &schemas.HealthcheckSiteTitleConfigResponse{
		SystemID:      model.SystemID,
		Name:          model.Name,
		Description:   model.Description,
		ExpectedTitle: model.ExpectedValue,
		Url:           model.Url,
		IsActive:      model.IsActive,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// models.HealthcheckConfigをschemas.HealthcheckEndpointConfigResponseに変換
func ConvertModelToHealthcheckEndpointConfigResponse(model *models.HealthcheckConfig) *schemas.HealthcheckEndpointConfigResponse {
	if model == nil {
		return nil
	}
	return &schemas.HealthcheckEndpointConfigResponse{
		SystemID:         model.SystemID,
		Name:             model.Name,
		Description:      model.Description,
		ExpectedResponse: model.ExpectedValue,
		Url:              model.Url,
		IsActive:         model.IsActive,
		CreatedAt:        model.CreatedAt,
		UpdatedAt:        model.UpdatedAt,
	}
}

// HealthcheckConfigの配列をHealthcheckConfigsResponseに変換
func ConvertHealthcheckConfigsToResponse(system models.System, configs []models.HealthcheckConfig) schemas.HealthcheckConfigsResponse {
	siteTitleConfigs := []schemas.HealthcheckSiteTitleConfigResponse{}
	endpointConfigs := []schemas.HealthcheckEndpointConfigResponse{}

	for _, config := range configs {
		switch config.ConfigType {
		case models.SiteTitle:
			siteTitleConfigs = append(siteTitleConfigs, schemas.HealthcheckSiteTitleConfigResponse{
				SystemID:      config.SystemID,
				Name:          config.Name,
				Description:   config.Description,
				ExpectedTitle: config.ExpectedValue,
				Url:           config.Url,
				IsActive:      config.IsActive,
				CreatedAt:     config.CreatedAt,
				UpdatedAt:     config.UpdatedAt,
			})
		case models.Endpoint:
			endpointConfigs = append(endpointConfigs, schemas.HealthcheckEndpointConfigResponse{
				SystemID:         config.SystemID,
				Name:             config.Name,
				Description:      config.Description,
				ExpectedResponse: config.ExpectedValue,
				Url:              config.Url,
				IsActive:         config.IsActive,
				CreatedAt:        config.CreatedAt,
				UpdatedAt:        config.UpdatedAt,
			})
		}
	}

	// システムモデルをレスポンススキーマに変換
	systemResponse := ConvertSystemModelToResponseSchema(&system)

	return schemas.HealthcheckConfigsResponse{
		SystemResponse:   *systemResponse,
		SiteTitleConfigs: siteTitleConfigs,
		EndpointConfigs:  endpointConfigs,
	}
}
