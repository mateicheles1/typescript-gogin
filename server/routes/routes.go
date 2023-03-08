package routes

import (
	"gogin/handlers"

	"github.com/gin-gonic/gin"
)

func HandleRoutes() {
	r := gin.Default()
	r.GET("/albums", handlers.Albums)

	r.GET("albums/:id", handlers.GetAlbumById)
	r.POST("albums/", handlers.CreateAlbumById)
	r.PATCH("albums/:id", handlers.UpdateAlbumById)
	r.DELETE("albums/:id", handlers.DeleteAlbumById)

	r.Run("localhost:8080")
}
