package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vitorbgouveia/go-web-api/internal/presentation/controllers"
)

// ConfigRoutes load http routes
func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("app/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.ShowBooks)
			books.GET("/:id", controllers.ShowBook)
			books.POST("/", controllers.CreateBook)
			books.PUT("/", controllers.UpdateBook)
			books.DELETE("/:id", controllers.DeleteBook)
		}
	}

	return router
}
