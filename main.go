package main

import (
	"log"
	"github.com/qiaqia/haha"
	"github.com/gin-gonic/gin"
)

func main() {
	rouder := gin.New()
	rouder.Use(Login(), gin.Logger(), gin.Recovery())
    rouder.GET("/ping", func(c *gin.Context) {

		random := c.Query("random")

		a := &haha.Haha{Name: "wangqiaqia"}
		a.SayHi()

        c.JSON(200, gin.H{
			"message": "pong",
			"random": random,
        })
    })
    rouder.Run() // listen and serve on 0.0.0.0:8080
}