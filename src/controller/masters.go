package controller

import (
	"log"
	http "net/http"

	crud "github.com/Dencyuman/logvista-server/src/crud"
	schemas "github.com/Dencyuman/logvista-server/src/schemas"
	utils "github.com/Dencyuman/logvista-server/src/utils"
	gin "github.com/gin-gonic/gin"
)

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
	errorTypesWithoutEmpty := utils.FilterEmptyStrings(errorTypes)
	c.JSON(http.StatusOK, errorTypesWithoutEmpty)
}

// @Summary ファイル名一覧取得
// @Description DB上に存在するファイル名一覧を取得する
// @Tags masters
// @Accept json
// @Produce json
// @Router /masters/files [get]
// @Success 200 {object} []string
// @Failure 500 {object} schemas.ErrorResponse
func (ctrl *AppController) GetFiles(c *gin.Context) {
	files, err := crud.FindFiles(ctrl.DB)
	if err != nil {
		log.Printf("Error finding files: %v\n", err)
		c.JSON(http.StatusInternalServerError, schemas.ErrorResponse{Message: "Internal Server Error"})
		return
	}
	filesWithoutEmpty := utils.FilterEmptyStrings(files)
	c.JSON(http.StatusOK, filesWithoutEmpty)
}

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
	levelsWithoutEmpty := utils.FilterEmptyStrings(levels)
	c.JSON(http.StatusOK, levelsWithoutEmpty)
}
