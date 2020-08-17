package main

import "github.com/gin-gonic/gin"

func hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "ok"})
}

func main() {
	r := gin.Default()
	r.GET("/cards", hello)

	r.Run(":1234")
}
