package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	err := r.Run() // 监听并在 0.0.0.0:8080 上启动服务
	if err != nil {
		fmt.Println("gin run server err:", err)
	}
}
