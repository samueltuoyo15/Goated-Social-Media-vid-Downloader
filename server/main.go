package main

import(
  "github.com/gin-gonic/gin"
  "path/filepath"
  )

func main(){
  router := gin.Default()
  
    distPath := filepath.Join("..", "client", "dist") 
    router.Static("/static", distPath) 

    router.NoRoute(func(c *gin.Context) {
        c.File(filepath.Join(distPath, "index.html")) 
    })
  
  router.SetTrustedProxies([]string{"127.0.0.1"})
  if err := router.Run(":10000"); err != nil{
    panic(err)
  }
}

