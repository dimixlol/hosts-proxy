package utils

import (
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
)

func getDefaultMiddleWares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		logging.JSONLogMiddleware(),
		gin.CustomRecovery(Recovery),
		cors.Default(),
	}
}

func NewEngine(middlewares ...gin.HandlerFunc) *fizz.Fizz {
	engine := gin.New()
	middlewares = append(middlewares, getDefaultMiddleWares()...)

	for _, middleware := range middlewares {
		engine.Use(middleware)
	}
	return fizz.NewFromEngine(engine)
}
