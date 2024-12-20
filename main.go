package main

import (
	"api-go-gin/controllers"
	"api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	//router
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Halaman index",
		})
	})
	router.GET("/api/posts", controllers.FindPosts)
	router.POST("/api/posts", controllers.StorePost)
	router.GET("/api/posts/:id", controllers.FindPostById)
	router.PUT("/api/posts/:id", controllers.UpdatePost)
	router.DELETE("/api/posts/:id", controllers.DeletePost)
	//end of router

	err := router.Run(":3000")
	if err != nil {
		return
	}
}
