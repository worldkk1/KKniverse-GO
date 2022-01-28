package donut

import "github.com/gin-gonic/gin"

func Routes(routes *gin.Engine) {
	router := routes.Group("/donut")
	{
		router.POST("/pointbalance", CitibankGetPoints)
	}
}
