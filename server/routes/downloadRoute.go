package routes

import (
  "net/http"
   "github.com/gin-gonic/gin"
   "palmdownload/controller"
  )
  
func fetchMetaData(c *gin.Context){
  videoURL := c.Query("url")
  
  if videoURL == "" {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Video url is required"})
    return 
  }
  
  metaData, err := controller.ExtractMetaData(videoURL)
  
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch meta data"})
    return 
  }
  
  c.JSON(http.StatusOK, metaData)
}

func setupRoutes(router, *gin.Engine){
  router.GET("/metadata", fetchMetaData)
}