package kaikong

import "github.com/gin-gonic/gin"

func Routes(routes *gin.Engine) {
	router := routes.Group("/kaikong/v1")
	{
		router.GET("/shop/live/available-lives", GetAvailableLives)
		router.GET("/shop/products", SearchProducts)
		router.POST("/shop/live/create", CreateLive)
		router.GET("/shop/lives", GetLiveList)
		router.GET("/shop/live/:live_id", GetLiveInfo)
		router.GET("/shop/live/:live_id/real-time-comment", GetLiveComments)
		router.POST("/shop/live/:live_id/comment", SendLiveMessage)
		router.POST("/shop/live/:live_id/send-invoices", SendAllInvoices)
	}
}
