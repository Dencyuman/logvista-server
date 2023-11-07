package controller

import (
	"log"
	http "net/http"
	"time"

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

// @Summary システム集計情報取得
// @Description DB上に存在するシステム別集計情報を取得する
// @Tags systems
// @Accept json
// @Produce json
// @Router /systems/summary [get]
// @Success 200 {object} []schemas.Summary
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetSystemSummary(c *gin.Context) {
	modelsSystems, err := crud.FindAllSystems(ctrl.DB)
	if err != nil {
		log.Printf("Error finding systems: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	var Summaries []schemas.Summary
	now := time.Now()
	for _, modelsSystem := range modelsSystems {
		summaryData, err := crud.FindSummaryData(ctrl.DB, &modelsSystem, 12, now)
		if err != nil {
			log.Printf("Error finding summary data: %v\n", err)
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
			return
		}
		latestLog, err := crud.FindLatestLog(ctrl.DB, modelsSystem.Name)
		if err != nil {
			log.Printf("Error finding latest log: %v\n", err)
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
			return
		}
		summary := converter.ConvertSystemModelAndSummaryDataToSchema(&modelsSystem, summaryData, latestLog)
		if summary == nil {
			log.Printf("Error converting system model and summary data to schema: %v\n", modelsSystem.ID)
			continue
		}
		Summaries = append(Summaries, *summary)
	}

	c.JSON(http.StatusOK, Summaries)
}
