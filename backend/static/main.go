package static

import (
	"embed"
	"io/fs"
)

var (
	//go:embed swagger-ui
	swaggerUIDir embed.FS
	//go:embed redoc
	redocFile embed.FS

	RedocAssets   fs.FS
	SwaggerAssets fs.FS
)

func init() {
	var err error
	SwaggerAssets, err = fs.Sub(swaggerUIDir, "swagger-ui")
	if err != nil {
		panic(err)
	}
	RedocAssets, err = fs.Sub(redocFile, "redoc")
	if err != nil {
		panic(err)
	}
}
