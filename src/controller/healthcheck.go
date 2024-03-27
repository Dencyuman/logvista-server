package controller

import (
	"fmt"
	"github.com/Dencyuman/logvista-server/src/background"
	"github.com/Dencyuman/logvista-server/src/converter"
	"github.com/Dencyuman/logvista-server/src/crud"
	"github.com/Dencyuman/logvista-server/src/schemas"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// @Summary ヘルスチェック設定取得用エンドポイント
// @Tags healthcheck
// @Description 200 システム別ヘルスチェック設定一覧の取得
// @Accept json
// @Produce json
// @Success 200 {object} []schemas.HealthcheckConfigsResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/ [get]
func (ctrl *AppController) GetHealthcheckConfigs(c *gin.Context) {
	systems, err := crud.FindAllSystems(ctrl.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	var healthcheckConfigsResponse []schemas.HealthcheckConfigsResponse
	for _, system := range systems {
		healthcheckConfigs, err := crud.FindHealthcheckConfigs(ctrl.DB, system.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
			return
		}
		healthcheckConfig := converter.ConvertHealthcheckConfigsToResponse(system, healthcheckConfigs)

		healthcheckConfigsResponse = append(healthcheckConfigsResponse, healthcheckConfig)
		fmt.Println(healthcheckConfigs)
	}

	c.JSON(http.StatusOK, healthcheckConfigsResponse)
}

// @Summary ヘルスチェック設定取得用エンドポイント(システム別)
// @Tags healthcheck
// @Description 200 systemIdで指定したシステムに紐づくヘルスチェック設定一覧の取得
// @Accept json
// @Produce json
// @Success 200 {object} schemas.HealthcheckConfigsResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/systems/{systemId} [get]
// @Param systemId path string true "システムid"
func (ctrl *AppController) GetSystemHealthcheckConfigs(c *gin.Context) {
	// パスパラメータからIDを取得する
	systemID := c.Param("systemId")

	// IDが提供されていない場合は、エラーメッセージを返す
	if systemID == "" {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "System ID is required"})
		return
	}

	// パスパラメータから取得したIDを使用してシステムを見つける
	modelsSystem, err := crud.FindSystemByID(ctrl.DB, systemID)
	if err != nil {
		log.Printf("Error finding system by id: %v\n", err)
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "System not found"})
		return
	}

	healthcheckConfigs, err := crud.FindHealthcheckConfigs(ctrl.DB, modelsSystem.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}
	healthcheckConfigsResponse := converter.ConvertHealthcheckConfigsToResponse(*modelsSystem, healthcheckConfigs)
	fmt.Println(healthcheckConfigs)

	c.JSON(http.StatusOK, healthcheckConfigsResponse)
}

// @Summary ヘルスチェック設定テスト用エンドポイント(SiteTitle)
// @Tags healthcheck
// @Description 200 設定通りにSiteTitleヘルスチェックを１回実行した結果を取得できる
// @Accept json
// @Produce json
// @Success 200 {object} schemas.TestHealthcheckSiteTitleConfigResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/site-title/test [post]
// @Param config body schemas.TestHealthcheckSiteTitleConfigBody false "SiteTitle用設定値"
func TestHealthcheckSiteTitleConfig(c *gin.Context) {
	var config schemas.TestHealthcheckSiteTitleConfigBody
	if err := c.ShouldBindJSON(&config); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	fetchedTitle, err := background.FetchPageTitle(config.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expected_title":     config.ExpectedTitle,
		"fetched_title":      fetchedTitle,
		"healthcheck_result": fetchedTitle == config.ExpectedTitle,
	})
}

// @Summary ヘルスチェック設定テスト用エンドポイント(Endpoint)
// @Tags healthcheck
// @Description 200 設定通りにEndpointヘルスチェックを１回実行した結果を取得できる
// @Accept json
// @Produce json
// @Success 200 {object} schemas.TestHealthcheckEndpointConfigResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/endpoint/test [post]
// @Param config body schemas.TestHealthcheckEndpointConfigBody false "Endpoint用設定値"
func TestHealthcheckEndpointConfig(c *gin.Context) {
	var config schemas.TestHealthcheckEndpointConfigBody
	if err := c.ShouldBindJSON(&config); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	response, err := background.FetchHealthcheckAPIResponseAsString(config.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expected_title":     config.ExpectedStatus,
		"fetched_title":      response,
		"healthcheck_result": response == config.ExpectedStatus,
	})
}

// @Summary ヘルスチェック設定用エンドポイント(SiteTitle)
// @Tags healthcheck
// @Description 200 SiteTitleヘルスチェックを設定する
// @Accept json
// @Produce json
// @Success 200 {object} schemas.HealthcheckSiteTitleConfigResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/site-title [post]
// @Param config body schemas.HealthcheckSiteTitleConfigBody false "SiteTitle用設定値"
func (ctrl *AppController) HealthcheckSiteTitleConfig(c *gin.Context) {
	var config schemas.HealthcheckSiteTitleConfigBody
	if err := c.ShouldBindJSON(&config); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	_, err := crud.FindSystemByID(ctrl.DB, config.SystemID)
	if err != nil {
		log.Printf("Error inserting log: %v\n", err)
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "System Not Found"})
		return
	}

	modelObject := converter.ConvertHealthcheckSiteTitleConfigBodyToModel(&config)

	if err := crud.InsertHealthcheck(ctrl.DB, modelObject); err != nil {
		log.Printf("Error inserting log: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	responseObject := converter.ConvertModelToHealthcheckSiteTitleConfigResponse(modelObject)

	c.JSON(http.StatusOK, responseObject)
}

// @Summary ヘルスチェック設定用エンドポイント(Endpoint)
// @Tags healthcheck
// @Description 200 Endpointヘルスチェックを設定する
// @Accept json
// @Produce json
// @Success 200 {object} schemas.HealthcheckEndpointConfigResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/endpoint [post]
// @Param config body schemas.HealthcheckEndpointConfigBody false "Endpoint用設定値"
func (ctrl *AppController) HealthcheckEndpointConfig(c *gin.Context) {
	var config schemas.HealthcheckEndpointConfigBody
	if err := c.ShouldBindJSON(&config); err != nil {
		log.Printf("Error binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Bad Request"})
		return
	}

	_, err := crud.FindSystemByID(ctrl.DB, config.SystemID)
	if err != nil {
		log.Printf("Error inserting log: %v\n", err)
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "System Not Found"})
		return
	}

	modelObject := converter.ConvertHealthcheckEndpointConfigBodyToModel(&config)

	if err := crud.InsertHealthcheck(ctrl.DB, modelObject); err != nil {
		log.Printf("Error inserting log: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}

	responseObject := converter.ConvertModelToHealthcheckEndpointConfigResponse(modelObject)

	c.JSON(http.StatusOK, responseObject)
}

// @Summary ヘルスチェックログ取得用エンドポイント
// @Tags healthcheck
// @Description 200 ヘルスチェックログ一覧の取得
// @Accept json
// @Produce json
// @Success 200 {object} []schemas.HealthcheckLogsListResponse
// @Failure 500 {object} schemas.ErrorResponse
// @Router /healthcheck/configs/{configId}/logs [get]
// @Param configId path string true "設定ID"
// @Param count query int false "取得ログデータ件数" default(10)
// @Param desc query bool false "降順フラグ" default(true)
func (ctrl *AppController) GetHealthcheckLogs(c *gin.Context) {
	configID := c.Param("configId")
	count, err := strconv.Atoi(c.Query("count"))
	desc, err := strconv.ParseBool(c.Query("desc"))
	if configID == "" {
		c.JSON(http.StatusBadRequest, schemas.ErrorResponse{Message: "Config ID is required"})
		return
	}
	config, err := crud.FindHealthcheckConfigByID(ctrl.DB, configID)
	if err != nil {
		c.JSON(http.StatusNotFound, schemas.ErrorResponse{Message: "Config Not Found"})
		return
	}
	healthcheckLogs, err := crud.FindHealthcheckLogs(ctrl.DB, configID, count, desc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	// 取得したConfigをレスポンスに変換
	healthcheckConfig := converter.ConvertHealthcheckConfigToResponse(config)

	// 取得したログをレスポンスに変換
	var healthcheckLogsResponse []schemas.HealthcheckLogsResponse
	for _, healthcheckLog := range healthcheckLogs {
		healthcheckLogSchema := converter.ConvertModelToHealthcheckLogsResponse(&healthcheckLog)
		healthcheckLogsResponse = append(healthcheckLogsResponse, *healthcheckLogSchema)
	}

	healthcheckLogsList := schemas.HealthcheckLogsListResponse{
		Config: *healthcheckConfig,
		Logs:   healthcheckLogsResponse,
	}

	c.JSON(http.StatusOK, healthcheckLogsList)
}
