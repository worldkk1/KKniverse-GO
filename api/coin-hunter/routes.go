package coin_hunter

import "github.com/gin-gonic/gin"

func Routes(routes *gin.Engine) {
	router := routes.Group("/coinhunter")
	{
		router.POST("/users/track", Braze)
	}
}
