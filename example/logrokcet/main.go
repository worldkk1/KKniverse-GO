// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
package main

import (
	"KK/KKniverse-GO/example/logrokcet/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.CreateBook)

	r.Run()
}
