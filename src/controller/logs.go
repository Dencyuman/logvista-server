package controller

import (
	"log"
	http "net/http"
	"time"

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
// @Router /logs/python-logvista [post]
// @Success 200 {object} []schemas.LogResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Param logs body []schemas.Log false "ログデータ"
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

// GetLogsのパラメータ型
type getLogsParams struct {
	Page        int       `form:"page,default=1" binding:"min=1"` // 現在のページ
	PageSize    int       `form:"pageSize,default=10" binding:"min=1"` // 1ページあたりのアイテム数
	StartDate   *time.Time `form:"startDate"` // 検索する日付の開始範囲
	EndDate     *time.Time `form:"endDate"` // 検索する日付の終了範囲
	LevelName   *string    `form:"levelName"` // ログレベル
	SystemName  *string    `form:"systemName"` // システム名
	ContainsMsg *string    `form:"containMsg"` // メッセージ内容を含むかどうか
	ExcType     *string    `form:"excType"` // エラーの種類
	ExcDetail   *string	   `form:"excDetail"` // エラーの詳細
	FileName    *string    `form:"fileName"` // ファイル名
	Lineno      *int       `form:"lineno"` // エラーが発生した行番号
}

// @Summary 取得ログ情報
// @Description 蓄積されているログ情報を取得する。クエリパラメータでフィルタリングが可能。
// @Tags logs
// @Accept json
// @Produce json
// @Router /logs/ [get]
// @Success 200 {array} schemas.PaginatedLogResponse
// @Failure 400 {object} schemas.ErrorResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Param page query int false "現在のページ" default(1) minimum(1)
// @Param pageSize query int false "1ページあたりのアイテム数" default(10) minimum(1)
// @Param startDate query string false "検索する日付の開始範囲 (形式: YYYY-MM-DDTHH:MM:SSZ)"
// @Param endDate query string false "検索する日付の終了範囲 (形式: YYYY-MM-DDTHH:MM:SSZ)"
// @Param levelName query string false "ログレベルでのフィルタ"
// @Param systemName query string false "システム名でのフィルタ"
// @Param containMsg query string false "メッセージ内容の部分一致フィルタ"
// @Param excType query string false "エラーの種類でのフィルタ"
// @Param excDetail query string false "エラーの詳細でのキーワード部分一致フィルタ"
// @Param fileName query string false "ファイル名でのフィルタ"
// @Param lineno query int false "エラーが発生した行番号でのフィルタ"
func (ctrl *AppController) GetLogs(c *gin.Context) {
	var filteredModelLogs []models.Log
	var params getLogsParams

	if err := c.ShouldBindQuery(&params); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// クエリパラメータをオプションに変換
	offset := (params.Page - 1) * params.PageSize
	opts := &crud.FindLogsOptions{
		Offset:      &offset,
		Limit:       &params.PageSize,
		StartDate:   params.StartDate,
		EndDate:     params.EndDate,
		LevelName:   params.LevelName,
		SystemName:  params.SystemName,
		ContainsMsg: params.ContainsMsg,
		ExcType:     params.ExcType,
		ExcDetail:   params.ExcDetail,
		FileName:    params.FileName,
		Lineno:      params.Lineno,
	}

	// オプションを指定してログを取得
	filteredModelLogs, err := crud.FindLogs(ctrl.DB, opts)
	if err != nil {
		log.Printf("Error finding logs: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	log.Printf("params: %v\n", params)
	log.Printf("Found logs count: %d\n", len(filteredModelLogs))

	// オプションを変更して全件取得
	opts.Offset = nil
	opts.Limit = nil
	allModelLogs, err := crud.FindLogs(ctrl.DB, opts)
	if err != nil {
		log.Printf("Error finding logs: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	// models.Logをschemas.Logに変換
	schemaLogs := make([]schemas.LogResponse, 0)
	for _, filteredModelLog := range filteredModelLogs {
		schemaLog := converter.ConvertLogModelToResponseSchema(&filteredModelLog)
		if schemaLog == nil {
			log.Printf("Error converting log model to schema: %v\n", filteredModelLog.ID)
			continue
		}
		schemaLogs = append(schemaLogs, *schemaLog)
	}

	// ページネーション用のレスポンスを作成
	paginatedResponse := schemas.PaginatedLogResponse{
		Total: len(allModelLogs),
		Page: params.Page,
		Limit: params.PageSize,
		TotalPages: len(allModelLogs) / params.PageSize + 1,
		Items: schemaLogs,
	}
	c.JSON(http.StatusOK, paginatedResponse)
}
