package controller

import (
	"log"
	http "net/http"

	converter "github.com/Dencyuman/logvista-server/src/converter"
	crud "github.com/Dencyuman/logvista-server/src/crud"
	models "github.com/Dencyuman/logvista-server/src/models"
	schemas "github.com/Dencyuman/logvista-server/src/schemas"
	gin "github.com/gin-gonic/gin"
)

// @Summary python-logvista用エンドポイント
// @Tags logs
// @Description json形式の配列で受け取ったログ情報を記録する
// @Accept json
// @Produce json
// @Param logs body []schemas.Log false "ログデータ"
// @Success 200 {object} []schemas.LogResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /logs/python-logvista [post]
func (ctrl *AppController) RecordLogs(c *gin.Context) {
	var schemaLogs []schemas.Log
	if err := c.ShouldBindJSON(&schemaLogs); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	// schemas.Logをmodels.Logに変換
	var modelLogs []models.Log
	for _, schemaLog := range schemaLogs {
		modelLog := converter.ConvertLogSchemaToModel(&schemaLog)
		if modelLog == nil {
			log.Printf("Error converting log schema to model: %v\n", schemaLog.ID)
			continue
		}
		modelLogs = append(modelLogs, *modelLog)
	}

	// データベースにログとそのトレースバックを挿入
	for _, modelLog := range modelLogs {
		if err := crud.InsertLogWithTracebacks(ctrl.DB, &modelLog); err != nil {
			log.Printf("Error inserting log: %v\n", err)
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
			return
		}
		log.Println("Inserted log:", modelLog.ID)
	}
	c.JSON(http.StatusOK, schemaLogs)
}

// @Summary log情報取得用エンドポイント
// @Tags logs
// @Description 蓄積されているログ情報の一覧を取得する
// @Accept json
// @Produce json
// @Success 200 {object} []schemas.LogResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Router /logs/ [get]
func (ctrl *AppController) GetLogs(c *gin.Context) {
	var modelLogs []models.Log

	opts := &crud.FindLogsOptions{
		Limit:  nil,
		Offset: nil,
	}

	modelLogs, err := crud.FindLogs(ctrl.DB, opts)
	if err != nil {
		log.Printf("Error finding logs: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	// models.Logをschemas.Logに変換
	schemaLogs := make([]schemas.LogResponse, 0)
	for _, modelLog := range modelLogs {
		schemaLog := converter.ConvertLogModelToResponseSchema(&modelLog)
		if schemaLog == nil {
			log.Printf("Error converting log model to schema: %v\n", modelLog.ID)
			continue
		}
		schemaLogs = append(schemaLogs, *schemaLog)
	}
	c.JSON(http.StatusOK, schemaLogs)
}
