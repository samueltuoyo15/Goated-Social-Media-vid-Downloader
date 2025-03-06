package main

import (
	"github.com/gin-gonic/gin"
	"path/filepath"
	"palmdownload/routes"
	"os"
)

func main() {
	router := gin.Default()

	routes.SetupRoutes(router)

	rootDir, _ := os.Getwd()
	distPath := filepath.Join(rootDir, "../client/dist")

	router.StaticFS("/", gin.Dir(distPath, false))

	router.NoRoute(func(c *gin.Context) {
		c.File(filepath.Join(distPath, "index.html"))
	})

	router.SetTrustedProxies(nil)

	if err := router.Run(":10000"); err != nil {
		panic(err)
	}
}