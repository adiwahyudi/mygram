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

	g.GET("", controller.BaseContoller)
	base := g.Group("/api/v1")
	{
		base.GET("/mygram", middleware.AuthMiddleware, userController.MyGram)
		auth := base.Group("/auth")
		{
			auth.POST("/register", userController.Register)
			auth.POST("/login", userController.Login)
		}
		socialMediaRoute := base.Group("/social_media", middleware.AuthMiddleware)
		{
			socialMediaRoute.GET("", socialMediaController.GetListSocialMedias)
			socialMediaRoute.GET("/:id", socialMediaController.GetOneSocialMediaByID)
			socialMediaRoute.POST("", socialMediaController.CreateSocialMedia)
			socialMediaRoute.PUT("/:id", socialMediaController.UpdateSocialMedia)
			socialMediaRoute.DELETE("/:id", socialMediaController.DeleteSocialMedia)

		}
		photoRoute := base.Group("/photo", middleware.AuthMiddleware)
		{
			photoRoute.GET("", photoController.GetListPhotos)
			photoRoute.GET("/:id", photoController.GetPhotoByID)
			photoRoute.POST("", photoController.CreatePhoto)
			photoRoute.PUT("/:id", photoController.UpdatePhoto)
			photoRoute.DELETE("/:id", photoController.DeletePhoto)
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
