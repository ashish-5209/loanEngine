package main

import (
	"loanEngine/common"
	"loanEngine/docs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

var (
	mainRouter *gin.Engine
)

func init() {
	mainRouter = gin.Default()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(mainRouter)
	// Ping test
	mainRouter.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Setup Swagger
	docs.SwaggerInfo.BasePath = "/api/v1"
	mainRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// mainRouter.Use(common.Middleware) // extract common meta info
	AddLoanRoutes(mainRouter)
}
func main() {
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatal("APP_PORT environment variable not set")
	}
	err := mainRouter.Run(":" + appPort)
	if err != nil {
		common.CriticalLogger.Fatal("Failed to start server: ", err)
	}
	common.Logger.Info("server started!!")
}
