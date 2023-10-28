package controller

import (
	"log"
	http "net/http"

	crud "github.com/Dencyuman/logvista-server/src/crud"
	schemas "github.com/Dencyuman/logvista-server/src/schemas"
	gin "github.com/gin-gonic/gin"
)

// @Summary ログレベル一覧取得
// @Description DB上に存在するログレベル一覧を取得する
// @Tags masters
// @Accept json
// @Produce json
// @Router /masters/levels [get]
// @Success 200 {object} []string
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetLevels(c *gin.Context) {
	levels, err := crud.FindLevels(ctrl.DB)
	if err != nil {
		log.Printf("Error finding levels: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, levels)
}

// @Summary システム名一覧取得
// @Description DB上に存在するシステム名一覧を取得する
// @Tags masters
// @Accept json
// @Produce json
// @Router /masters/systems [get]
// @Success 200 {object} []string
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetSystems(c *gin.Context) {
	systems, err := crud.FindSystems(ctrl.DB)
	if err != nil {
		log.Printf("Error finding systems: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, systems)
}

// @Summary エラー型名一覧取得
// @Description DB上に存在するエラー型名一覧を取得する
// @Tags masters
// @Accept json
// @Produce json
// @Router /masters/error-types [get]
// @Success 200 {object} []string
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetErrorTypes(c *gin.Context) {
	errorTypes, err := crud.FindErrorTypes(ctrl.DB)
	if err != nil {
		log.Printf("Error finding error types: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, errorTypes)
}
