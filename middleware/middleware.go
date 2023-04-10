package middleware

import (
	"mygram/helper"
	"mygram/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")

	if auth == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
			Err: model.ErrorNotAuthorized.Err,
		})
		return
	}
	token := strings.Split(auth, " ")[1]

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
			Err: model.ErrorNotAuthorized.Err,
		})
		return
	}

	jwtToken, err := helper.VerifyToken(token)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.Set("user_id", claims["user_id"])

	ctx.Next()
}
