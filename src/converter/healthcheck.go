package converter

import (
	"fmt"
	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"
	"strconv"
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
		Timespan:      body.Timespan,
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
		ExpectedValue: fmt.Sprintf("%d", body.ExpectedStatus),
		Url:           body.Url,
		Timespan:      body.Timespan,
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
		Timespan:      model.Timespan,
		CreatedAt:     model.CreatedAt,
		UpdatedAt:     model.UpdatedAt,
	}
}

// models.HealthcheckConfigをschemas.HealthcheckEndpointConfigResponseに変換
func ConvertModelToHealthcheckEndpointConfigResponse(model *models.HealthcheckConfig) *schemas.HealthcheckEndpointConfigResponse {
	if model == nil {
		return nil
	}
	expectedStatus, _ := strconv.Atoi(model.ExpectedValue) // Convert ExpectedValue string to int
	return &schemas.HealthcheckEndpointConfigResponse{
		SystemID:       model.SystemID,
		Name:           model.Name,
		Description:    model.Description,
		ExpectedStatus: expectedStatus,
		Url:            model.Url,
		Timespan:       model.Timespan,
		CreatedAt:      model.CreatedAt,
		UpdatedAt:      model.UpdatedAt,
	}
}
