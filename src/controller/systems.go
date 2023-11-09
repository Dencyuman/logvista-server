package controller

import (
	"log"
	http "net/http"
	"sort"
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

	// systemResponses を CreatedAt フィールドでソートする
	sort.Slice(systemResponses, func(i, j int) bool {
		return systemResponses[i].CreatedAt.Before(systemResponses[j].CreatedAt)
	})

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
	roundedTime := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location())
	for _, modelsSystem := range modelsSystems {
		summaryData, err := crud.FindSummaryData(ctrl.DB, &modelsSystem, 12, roundedTime)
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

	// Summaries を CreatedAt フィールドでソートする
	sort.Slice(Summaries, func(i, j int) bool {
		return Summaries[i].CreatedAt.Before(Summaries[j].CreatedAt)
	})

	c.JSON(http.StatusOK, Summaries)
}

// @Summary システム更新
// @Description DB上に存在するシステムを更新する
// @Tags systems
// @Accept json
// @Produce json
// @Router /systems/ [put]
// @Param ID query string true "システムID"
// @Param system body schemas.SystemRequest true "Update System Request"
// @Success 200 {string} string "OK"
// @Failure 400 {object} schemas.ErrorResponse
func (ctrl *AppController) UpdateSystem(c *gin.Context) {
	// クエリパラメータからIDを取得する
	systemID := c.Query("ID")

	// IDが提供されていない場合は、エラーメッセージを返す
	if systemID == "" {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "System ID is required"})
		return
	}

	var systemRequest schemas.SystemRequest
	if err := c.ShouldBindJSON(&systemRequest); err != nil {
		log.Printf("Error binding system request: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	// クエリパラメータから取得したIDを使用してシステムを見つける
	modelsSystem, err := crud.FindSystemByID(ctrl.DB, systemID)
	if err != nil {
		log.Printf("Error finding system by id: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	// JSONリクエストボディから受け取ったデータでシステムを更新する
	modelsSystem.Category = systemRequest.Category
	modelsSystem.UpdatedAt = time.Now()
	if err := crud.UpdateSystem(ctrl.DB, modelsSystem); err != nil {
		log.Printf("Error updating system: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, "OK")
}

// @Summary ログレベル割合一覧取得
// @Description 指定したシステムのログレベルの割合を取得する
// @Tags systems
// @Accept json
// @Produce json
// @Router /systems/{id}/level-counts [get]
// @Param id path string true "システムID"
// @Success 200 {object} []schemas.LevelCountResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetLevelCounts(c *gin.Context) {}