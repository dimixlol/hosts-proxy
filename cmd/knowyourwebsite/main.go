package main

import (
	"context"
	"embed"
	"fmt"
	"github.com/dimixlol/knowyourwebsite/config"
	"github.com/dimixlol/knowyourwebsite/domains/persister"
	"github.com/dimixlol/knowyourwebsite/domains/requester"
	"github.com/dimixlol/knowyourwebsite/logging"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
	"io/fs"
	"net/http"
	"time"
)

// github.com/elazarl/goproxy
//var customTransport = http.DefaultTransport
//
//var urls = map[string]string{
//	"foo": "199.188.205.68",
//	"bar": "199.188.205.65",
//}

//

//go:embed swagger-ui
var swaggerUIDir embed.FS

func registerRoutes(f *fizz.Fizz) {
	info := &openapi.Info{
		Title:       "Knowyourwebsite API",
		Description: "Manage mapping between domains and IPs",
		Version:     "0.0.1",
	}
	f.GET("/api/v1/swagger.json", nil, f.OpenAPI(info, "json"))
	f.GET("/api/v1/swagger.yaml", nil, f.OpenAPI(info, "yaml"))
}

func main() {
	config.CreateConfiguration("config.json")
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(logging.JSONLogMiddleware())
	engine.Use(gin.Recovery())
	//engine.Use(gin.CustomRecovery(response.Recovery))
	engine.Use(cors.Default())

	//engine.Use(...) // register global middlewares
	f := fizz.NewFromEngine(engine)
	f.Use(logging.JSONLogMiddleware())
	swaggerAssets, fsErr := fs.Sub(swaggerUIDir, "swagger-ui")
	if fsErr != nil {
		panic(fsErr)
	}
	engine.StaticFS("/api/swagger", http.FS(swaggerAssets))
	engine.GET("/", func(c *gin.Context) { c.Redirect(http.StatusMovedPermanently, "/api/swagger") })

	ctx := context.Background()
	const APIVersion = "v1"
	persister.NewPersister(ctx, f)
	registerRoutes(f)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Host, config.Configuration.Port),
		Handler: f,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			fmt.Println("exited with error:", err)
		}
	}()

	proxy := gin.New()
	proxy.Use(logging.JSONLogMiddleware())
	proxy.Use(gin.Recovery())
	req := requester.NewRequester(ctx)
	proxy.Any("/*any", req)
	srv2 := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Configuration.Host, "8081"),
		Handler: proxy,
	}
	go func() {
		err := srv2.ListenAndServe()
		if err != nil {
			fmt.Println("exited with error:", err)
		}
	}()
	time.Sleep(100 * time.Minute)
}
