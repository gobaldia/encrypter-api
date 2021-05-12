package main

import (
	"github.com/dvwright/xss-mw"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gobaldia/encrypter-api/handler"
)

func main() {
	r := gin.New()
	r.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Content-Type"}
	r.Use(cors.New(corsConfig))

	var xssMdlwr xss.XssMw
	r.Use(xssMdlwr.RemoveXss())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/encrypt", handler.EncryptQuote)

	r.Run(":3000")
}