package controller

import (
	"mygram/model"
	"mygram/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}

// Register godoc
//
//	@Summary		Register User
//	@Description	Sign up for user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserRegisterRequest	true	"User request is required"
//	@Success		201		{object}	model.UserRegisterResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Router			/auth/register [post]
func (uc *UserController) Register(ctx *gin.Context) {
	newUser := model.UserRegisterRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	res, err := uc.UserService.Add(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)

}

// Login godoc
//
//	@Summary		Login User
//	@Description	Sign in for user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserLoginRequest	true	"User request is required"
//	@Success		200		{object}	model.UserLoginResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Router			/auth/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	newUser := model.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	res, err := uc.UserService.Login(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.MyError{
			Err: model.ErrorInvalidEmailOrPassword.Err,
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}
