package main

import (
	"EverDownload/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(router)

	rootDir, _ := os.Getwd()
	distPath := filepath.Join(rootDir, "client/dist")

	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/metadata" || c.Request.URL.Path == "/proxy" {
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