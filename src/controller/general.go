package controller

import (
	gin "github.com/gin-gonic/gin"
)

// @Summary ヘルスチェック用エンドポイント
// @Tags general
// @Description 200 OKが返ってくれば起動済み
// @Accept json
// @Produce json
// @Success 200 {object} schemas.ResponseMessage
// @Router /healthcheck [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "OK",
	})
}
