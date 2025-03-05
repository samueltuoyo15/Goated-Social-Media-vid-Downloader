package main

import(
  "github.com/gin-gonic/gin"
  "palmdownload/routes"
  )

func main(){
  router := gin.Default()
  
  routes.SetupRoutes(router)
  router.Static("/static", "../client/dist")
  port := ":10000"
  
  if err := router.Run(port) err != nil{
    panic(err)
  }
}