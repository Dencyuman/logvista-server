package controller

import (
	"log"
	http "net/http"
	"sort"
	"strconv"
	"time"

	"github.com/Dencyuman/logvista-server/src/converter"
	crud "github.com/Dencyuman/logvista-server/src/crud"
	models "github.com/Dencyuman/logvista-server/src/models"
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
// @Param id query string false "システムID：指定しない場合は全てのシステムを取得"
// @Param timeSpan query int false "集計時間スパン（秒）: 10秒刻みで指定可能" minimum(10) default(3600)
// @Param dataCount query int false "取得データ個数" minimum(1) default(12)
// @Success 200 {object} []schemas.Summary
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetSystemSummary(c *gin.Context) {
	timeSpanParam := c.DefaultQuery("timeSpan", "3600") // デフォルトを1時間とする
	systemID := c.Query("id")                           // オプショナルのシステムID
	dataCountParam := c.DefaultQuery("dataCount", "12") // デフォルトを12個とする
	timeSpan, err := strconv.Atoi(timeSpanParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Invalid timeSpan parameter"})
		return
	}
	if timeSpan < 10 {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "timeSpan must be at least 10 seconds"})
		return
	}
	dataCount, err := strconv.Atoi(dataCountParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Invalid dataCount parameter"})
		return
	}

	now := time.Now()
	var roundedTime time.Time
	if timeSpan >= 60 {
		// timeSpanが60秒以上の場合は分または時を繰り上げる
		roundedMinute := now.Minute() + timeSpan/60 - now.Minute()%(timeSpan/60)
		if roundedMinute >= 60 {
			roundedMinute = 0
			now = now.Add(time.Hour)
		}
		roundedTime = time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), roundedMinute, 0, 0, now.Location())
	} else {
		// timeSpanが60秒未満の場合は秒を繰り上げる
		extraSeconds := timeSpan - now.Second()%timeSpan
		if now.Second()%timeSpan == 0 { // 既にtimeSpanの倍数の秒数の場合は繰り上げない
			extraSeconds = 0
		}
		roundedTime = now.Add(time.Duration(extraSeconds) * time.Second)
	}

	var Summaries []schemas.Summary
	var modelsSystems []models.System
	if systemID != "" {
		log.Printf("システム指定")
		system, err := crud.FindSystemByID(ctrl.DB, systemID)
		if err != nil {
			log.Printf("Error finding system by ID: %v\n", err)
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
			return
		}
		// appendを使用してスライスに追加する際は、ポインタの指す値を使用する
		modelsSystems = append(modelsSystems, *system) // ポインタをデリファレンスして値を取得
	} else {
		log.Printf("システム未指定")
		// システムIDが指定されていない場合は全てのシステムを取得
		modelsSystems, err = crud.FindAllSystems(ctrl.DB)
		if err != nil {
			log.Printf("Error finding systems: %v\n", err)
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
			return
		}
	}

	for _, modelsSystem := range modelsSystems {
		summaryData, err := crud.FindSummaryData(ctrl.DB, &modelsSystem, timeSpan, roundedTime, dataCount)
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
