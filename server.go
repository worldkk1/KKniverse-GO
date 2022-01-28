package main

import (
	coin_hunter "KK/KKniverse-GO/api/coin-hunter"
	demo_content_indexer "KK/KKniverse-GO/api/demo-content-indexer"
	"KK/KKniverse-GO/api/donut"
	"KK/KKniverse-GO/api/kaikong"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	kaikong.Routes(r)
	demo_content_indexer.Routes(r)
	coin_hunter.Routes(r)
	donut.Routes(r)

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8082")
}
