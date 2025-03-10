package routes

import (
	"EverDownload/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/metadata", fetchMetaData)
	router.GET("/proxy", controller.ProxyRequest)
}

func fetchMetaData(c *gin.Context) {
	videoURL := c.Query("url")
	if videoURL == "" {
		c.JSON(400, gin.H{"error": "Video URL is required"})
		return
	}

	metaData, err := controller.ExtractMetaData(videoURL)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch metadata"})
		return
	}

	metaData.Thumbnail = "/proxy?url=" + url.QueryEscape(metaData.Thumbnail)
	for i, link := range metaData.Links {
		metaData.Links[i].Link = "/proxy?url=" + url.QueryEscape(link.Link)
	}

	c.JSON(200, metaData)
}