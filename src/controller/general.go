package controller

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Logvistaヘルスチェック用エンドポイント
// @Tags general
// @Description 200 OKが返ってくれば起動済み
// @Accept json
// @Produce json
// @Success 200 {object} schemas.ResponseMessage
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
