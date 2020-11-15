package main

import (
	"net/http"
	"tokentaskV1/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.StaticFile("/", "disc/index.html") //指定静态页面
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//增加注册接口
	r.POST("/register", routes.Register)
	r.POST("/login", routes.Login)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
