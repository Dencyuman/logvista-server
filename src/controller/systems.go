package controller

import (
	"log"
	http "net/http"

	"github.com/Dencyuman/logvista-server/src/converter"
	crud "github.com/Dencyuman/logvista-server/src/crud"
	schemas "github.com/Dencyuman/logvista-server/src/schemas"
	gin "github.com/gin-gonic/gin"
)

// @Summary システム一覧取得
// @Description DB上に存在するシステム一覧を取得する
// @Tags systems
// @Accept json
// @Produce json
// @Router /systems/ [get]
// @Success 200 {object} []schemas.SystemResponse
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetSystems(c *gin.Context) {
	modelsSystems, err := crud.FindAllSystems(ctrl.DB)
	if err != nil {
		log.Printf("Error finding systems: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	var systemResponses []schemas.SystemResponse
	for _, modelsSystem := range modelsSystems {
		systemResponse := converter.ConvertSystemModelToResponseSchema(&modelsSystem)
		if systemResponse == nil {
			log.Printf("Error converting system model to response schema: %v\n", modelsSystem.ID)
			continue
		}
		systemResponses = append(systemResponses, *systemResponse)
	}

	c.JSON(http.StatusOK, systemResponses)
}