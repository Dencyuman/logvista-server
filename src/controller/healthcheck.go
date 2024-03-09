package controller

import (
	"github.com/Dencyuman/logvista-server/src/converter"
	"github.com/Dencyuman/logvista-server/src/crud"
	"github.com/Dencyuman/logvista-server/src/schemas"
	"github.com/Dencyuman/logvista-server/src/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

	fetchedTitle, err := utils.FetchPageTitle(config.Url)
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

	fetchedStatus, err := utils.FetchStatusCode(config.Url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"expected_title":     config.ExpectedStatus,
		"fetched_title":      fetchedStatus,
		"healthcheck_result": fetchedStatus == config.ExpectedStatus,
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
