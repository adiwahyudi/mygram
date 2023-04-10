package routes

import (
	"mygram/controller"
	"mygram/middleware"
	"mygram/repository"
	"mygram/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Routes(g *gin.Engine, db *gorm.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(*userService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controller.NewSocialMediaController(*socialMediaService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controller.NewPhotoController(*photoService)

	commentRepository := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepository, photoRepository)
	commentController := controller.NewCommentController(*commentService)

	base := g.Group("/api/v1")
	{
		base.GET("", controller.WelcomeController)
		auth := base.Group("/auth")
		{
			auth.POST("/register", userController.Register)
			auth.POST("/login", userController.Login)
		}
		socialMediaRoute := base.Group("/social_media", middleware.AuthMiddleware)
		{
			socialMediaRoute.GET("", socialMediaController.GetList)
			socialMediaRoute.GET("/:id", socialMediaController.GetByID)
			socialMediaRoute.POST("", socialMediaController.Create)
			socialMediaRoute.PUT("/:id", socialMediaController.Update)
			socialMediaRoute.DELETE("/:id", socialMediaController.Delete)

		}
		photoRoute := base.Group("/photo", middleware.AuthMiddleware)
		{
			photoRoute.GET("", photoController.GetList)
			photoRoute.GET("/:id", photoController.GetByID)
			photoRoute.POST("", photoController.Create)
			photoRoute.PUT("/:id", photoController.Update)
			photoRoute.DELETE("/:id", photoController.Delete)
		}

		commentRoute := base.Group("/comment", middleware.AuthMiddleware)
		{
			commentRoute.GET("", commentController.GetListComments)
			commentRoute.POST("/:photo_id", commentController.CreateCommentByPhotoID)
			commentRoute.GET("/:id", commentController.GetOneCommentsByID)
			commentRoute.PUT("/:id", commentController.UpdateComment)
			commentRoute.DELETE("/:id", commentController.DeleteComment)
		}
	}

}
