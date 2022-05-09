package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	println("Hello, world.")

	// create string variable
	var str string = "Hello"

	// if str is "Hello" then print "World!"
	if str == "Hello" {
		println("World!")
	}

	// print current time in unix
	println("Current time:", time.Now().Unix())

	// create gin web server
	router := gin.Default()

	// create root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"hello": "world"})
	})

	// route /ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	// start web server
	router.Run(":8080")
}
