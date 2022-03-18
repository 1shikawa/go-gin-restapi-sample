package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"massage": "Hello World",
		})
	})

	router.Run(":3000" )
}

func hello() string {
	return "Hello Golang"
}
