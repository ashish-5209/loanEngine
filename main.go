package main

import (
	"loanEngine/common"

	"github.com/gin-gonic/gin"
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
