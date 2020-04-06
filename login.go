package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func Login() gin.HandlerFunc {
	return func(c *gin.Context){
		log.Printf("Login checked")
		c.Next()
		log.Printf("after")
	}
}