package routes

import (
   "github.com/gin-gonic/gin"
   "EverDownload/controller"
  )
  
func fetchMetaData(c *gin.Context){
  videoURL := c.Query("url")
  
  if videoURL == "" {
    c.JSON(400, gin.H{"error": "Video url is required"})
    return 
  }
  
  metaData, err := controller.ExtractMetaData(videoURL)
  
  if err != nil {
    c.JSON(500, gin.H{"error": "Failed to fetch meta data"})
    return 
  }
  
  c.JSON(200, metaData)
}

func SetupRoutes(router *gin.Engine){
  router.GET("/metadata", fetchMetaData)
}