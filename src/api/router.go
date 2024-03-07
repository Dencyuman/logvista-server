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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

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
			masterGroup.GET("/error-types", appController.GetErrorTypes)
			masterGroup.GET("/files", appController.GetFiles)
			masterGroup.GET("/levels", appController.GetLevels)
			masterGroup.GET("/systems", appController.GetSystems)
		}
		systemsGroup := v1.Group("/systems")
		{
			systemsGroup.GET("/", appController.GetSystems)
			systemsGroup.GET("/summary", appController.GetSystemSummary)
			systemsGroup.PUT("/:systemName", appController.UpdateSystem)
			systemsGroup.DELETE("/:systemId", appController.DeleteSystem)
		}
	}

	router.Static("/static", "./static")
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	router.GET("/assets/*any", func(c *gin.Context) {
		filePath := c.Param("any")
		c.File("./static/assets/" + filePath)
	})
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	return router
}
