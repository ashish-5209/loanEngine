package main

import (
	"loanEngine/common"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	// mainRouter.Use(common.Middleware) // extract common meta info
	AddLoanRoutes(mainRouter)
}
func main() {
	err := mainRouter.Run(":5002")
	if err != nil {
		common.CriticalLogger.Fatal("Failed to start server: ", err)
	}
	common.Logger.Info("server started!!")
}
