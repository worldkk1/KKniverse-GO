package donut

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CitibankGetPoints(c *gin.Context) {
	// c.JSON(http.StatusOK, gin.H{
	// 	"code":            200,
	// 	"message":         "Success",
	// 	"detail":          "Successful operation",
	// 	"requestid":       "1643046180538",
	// 	"pointbalance":    19,
	// 	"refaisrequestid": "20220124_03",
	// })
	c.JSON(http.StatusBadRequest, gin.H{
		"code":            400,
		"message":         "inactive",
		"detail":          "Successful operation",
		"requestid":       "1643046180538",
		"pointbalance":    19,
		"refaisrequestid": "20220124_03",
	})
}
