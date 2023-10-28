package api

import (
	docs "github.com/Dencyuman/logvista-server/docs"
	controller "github.com/Dencyuman/logvista-server/src/controller"
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	gorm "gorm.io/gorm"
)

// SetupRouter is a function to initialize gin router
func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())

	appController := &controller.AppController{DB: db}

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		generalGroup := v1.Group("/")
		{
			generalGroup.GET("/healthcheck", controller.HealthCheck)
		}
		logsGroup := v1.Group("/logs")
		{
			logsGroup.POST("/python-logvista", appController.RecordLogs)
			logsGroup.GET("/", appController.GetLogs)
		}
		masterGroup := v1.Group("/masters")
		{
			masterGroup.GET("/levels", appController.GetLevels)
			masterGroup.GET("/systems", appController.GetSystems)
			masterGroup.GET("/error-types", appController.GetErrorTypes)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
