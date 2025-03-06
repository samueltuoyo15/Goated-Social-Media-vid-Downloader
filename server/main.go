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
	distPath := filepath.Join(rootDir, "client/dist")

  router.Use(func(c *gin.Context) {

		if c.Request.URL.Path == "/metadata" {

			c.Next()
			return
		}

		filePath := filepath.Join(distPath, c.Request.URL.Path)
		if _, err := os.Stat(filePath); err == nil {
			c.File(filePath)
		} else {
			c.File(filepath.Join(distPath, "index.html"))
		}
	})

	router.SetTrustedProxies(nil)

	if err := router.Run(":10000"); err != nil {
		panic(err)
	}
}