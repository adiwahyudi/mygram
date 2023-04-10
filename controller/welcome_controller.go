package controller

import (
	"mygram/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeController(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.ResponseMyInformation{
		Name:      "Adi Wahyudi",
		Github:    "https://github.com/adiwahyudi",
		LinkendIn: "https://linkedin.com/in/i-wayan-adi-wahyudi",
		Discord:   "Adi Wahyudi#4674",
	})
}
