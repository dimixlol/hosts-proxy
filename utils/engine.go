package utils

import (
	"github.com/dimixlol/hosts-proxy/config"
	"github.com/dimixlol/hosts-proxy/logging"
	"github.com/dimixlol/hosts-proxy/static"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"net/http"
)

const (
	SwaggerEndpoint = "/api/swagger"
)

func getDefaultMiddleWares() []gin.HandlerFunc {
	return []gin.HandlerFunc{
		logging.JSONLogMiddleware(),
		gin.CustomRecovery(Recovery),
		cors.New(*config.Configuration.CORS),
	}
}

func NewApiEngine(openApiConstructor func() *openapi.Info, middlewares ...gin.HandlerFunc) *fizz.Fizz {
	engine := NewEngine(middlewares...)

	openApiInfo := openApiConstructor()
	engine.GET(SwaggerEndpoint+".json", nil, engine.OpenAPI(openApiInfo, "json"))
	engine.GET(SwaggerEndpoint+".yaml", nil, engine.OpenAPI(openApiInfo, "yaml"))

	engine.Engine().StaticFS("/api/swagger", http.FS(static.SwaggerAssets))
	engine.Engine().StaticFS("/api/redoc", http.FS(static.RedocAssets))

	//engine.GET("/swagger", nil, func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, SwaggerEndpoint) })

	return engine
}

func NewEngine(middlewares ...gin.HandlerFunc) *fizz.Fizz {
	engine := gin.New()
	middlewares = append(middlewares, getDefaultMiddleWares()...)

	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	return fizz.NewFromEngine(engine)
}
