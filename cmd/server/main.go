package main

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/murilommen/rocketseat-api-project/internal/handlers"
	"github.com/murilommen/rocketseat-api-project/internal/storage"
)

func main() {
	router := gin.Default()

	userStorage := storage.NewUserStorage()
	userHandler := handlers.NewUserHandler(userStorage)

	userRoutes := router.Group("/api/users") 
	{
		userRoutes.GET("", userHandler.FindAll)
		userRoutes.GET("/:id", userHandler.FindById)
		userRoutes.POST("", userHandler.Insert)
		userRoutes.PUT("/:id", userHandler.Update)
		userRoutes.DELETE("/:id", userHandler.Delete)
		
	}

	err := router.Run(":8080")
	if err != nil {
		slog.Error("Error spinning up server")
	}
}