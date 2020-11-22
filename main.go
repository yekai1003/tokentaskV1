package main

import (
	"net/http"
	"tokentaskV1/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-session/gin-session"
)

func main() {
	r := gin.Default()
	r.Use(ginsession.New())                   //使用session中间件
	r.StaticFile("/", "dist/index.html")      //指定静态页面
	r.StaticFile("/index", "dist/index.html") //指定静态页面
	r.Static("static", "dist/static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//增加注册接口
	r.POST("/register", routes.Register)
	//登陆接口
	r.POST("/login", routes.Login)
	//挖矿接口
	r.POST("/mint", routes.Mint)
	//发布任务接口
	r.POST("/issue", routes.Issue)
	//任务修改接口
	r.POST("/update", routes.Modify)

	r.Run(":9090") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
