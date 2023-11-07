package converter

import (
	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"
)

// schemas.Systemをmodels.Systemに変換
func ConvertSystemSchemaToModel(systemSchema *schemas.System) *models.System {
	if systemSchema == nil {
		return nil
	}

	systemModel := &models.System{
		Name:     systemSchema.Name,
		Category: systemSchema.Category,
	}

	return systemModel
}

// models.Systemをschemas.SystemResponseに変換
func ConvertSystemModelToResponseSchema(systemModel *models.System) *schemas.SystemResponse {
	if systemModel == nil {
		return nil
	}

	systemResponse := &schemas.SystemResponse{
		ID:        systemModel.ID,
		System: schemas.System{
			Name:     systemModel.Name,
			Category: systemModel.Category,
		},
		CreatedAt: systemModel.CreatedAt,
		UpdatedAt: systemModel.UpdatedAt,
	}

	return systemResponse
}