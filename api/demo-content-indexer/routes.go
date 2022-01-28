package demo_content_indexer

import "github.com/gin-gonic/gin"

func Routes(routes *gin.Engine) {
	router := routes.Group("/contentindexer/v1")
	{
		router.GET("customsearch/v1", Search)
	}
}
