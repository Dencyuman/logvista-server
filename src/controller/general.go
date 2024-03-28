package controller

import (
	gin "github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Logvistaヘルスチェック用エンドポイント
// @Tags general
// @Description 200 OKが返ってくれば起動済み
// @Accept json
// @Produce plain
// @Success 200 {string} string "Logvista API Server is running!"
// @Router / [get]
func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Logvista API Server is running!")
}
