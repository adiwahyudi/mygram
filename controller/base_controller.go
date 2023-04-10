package controller

import (
	"mygram/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BaseContoller(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: model.ResponseMyInformation{
			About:     "Final Project - Hacktiv8 DTSFGA - Scalable Web Service with Golang",
			Name:      "Adi Wahyudi",
			Github:    "https://github.com/adiwahyudi",
			LinkendIn: "https://linkedin.com/in/i-wayan-adi-wahyudi",
			Discord:   "Adi Wahyudi#4674",
		},
	})
}
