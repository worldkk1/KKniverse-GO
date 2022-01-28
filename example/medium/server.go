// https://phayao.medium.com/%E0%B8%A1%E0%B8%B2%E0%B8%AA%E0%B8%A3%E0%B9%89%E0%B8%B2%E0%B8%87-restful-web-service-%E0%B8%94%E0%B9%89%E0%B8%A7%E0%B8%A2-gin-go-web-framework-c1a619e0d90a
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomHandler struct{}

func (h *CustomHandler) GetAllCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	h := CustomHandler{}

	r.GET("/customers", h.GetAllCustomers)

	return r
}

func main() {
	r := setupRouter()

	r.Run()
}
