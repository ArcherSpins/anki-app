package routers

import (
	"anki-project/controllers"
	"anki-project/middleware"
	"anki-project/repository"
	"anki-project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	cardRepo := repository.NewCardRepository(db)
	cardService := services.NewCardService(cardRepo)
	cardController := controllers.NewCardsController(cardService)

	userRoutes := r.Group("/")
	{
		userRoutes.POST("/login", userController.Login)
		userRoutes.POST("/register", userController.Register)
	}

	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.AuthMiddleware())
	{
		authRoutes.GET("/users", nil)
		authRoutes.GET("/users/:id", nil)
		authRoutes.GET("/cards", cardController.GetListOfCards)
		authRoutes.GET("/cards/:id", cardController.GetCard)
		authRoutes.POST("/cards/create", cardController.CreateCard)
		authRoutes.PUT("/cards/edit", cardController.EditCard)
		authRoutes.DELETE("/cards/:id", cardController.DeleteCard)
	}
}
