package main

import (
	"mygram/database"
	_ "mygram/docs"
	"mygram/routes"
	"os"

	"github.com/gin-gonic/gin"

	_ "github.com/joho/godotenv/autoload"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title							Mygram API
// @version							1.0
// @description						Final Project for Scalable Web Service with Golang - Batch 1, DTS-FGA.
// @schemes							https
// @host							https://mygram-api.up.railway.app
// @BasePath  						/api/v1
// @accept							json
// @produce							json
// @securityDefinitions.apikey		Bearer
// @in								header
// @name							Authorization
func main() {
	g := gin.Default()

	database.StartDB()
	db := database.GetDB()

	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.Routes(g, db)

	g.Run(":" + os.Getenv("PORT"))
}
