package router

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/status")
	}
	{
		bucket := v1.Group("/bucket")
		bucket.GET("/users")
		bucket.GET("/users/:id")
		bucket.POST("/users")
	}
	{
		azure := v1.Group("/blob")
		azure.GET("/users")
		azure.GET("/users/:id")
		azure.POST("/users")
	}
	r.Run(":9000")
}
