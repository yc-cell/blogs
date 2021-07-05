package main

import (
	"embed"
	"mizuki/framework/core-kit/init/initkit"
	"mizuki/framework/core-kit/service/restkit/router"
)

var swaggerAssets embed.FS

func main() {
	initkit.LoadConfig()
	router.SwaggerAssets = swaggerAssets
}
