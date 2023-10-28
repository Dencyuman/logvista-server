package controller

import (
	"log"
	http "net/http"

	crud "github.com/Dencyuman/logvista-server/src/crud"
	schemas "github.com/Dencyuman/logvista-server/src/schemas"
	gin "github.com/gin-gonic/gin"
)

// @Summary システム名一覧取得
// @Description システム名一覧を取得する
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
