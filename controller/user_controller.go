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
//	@Success		201		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Router			/auth/register [post]
func (uc *UserController) Register(ctx *gin.Context) {
	newUser := model.UserRegisterRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(newUser)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	res, err := uc.UserService.Add(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusCreated,
			Message: http.StatusText(http.StatusCreated),
		},
		Data: res,
	})
	return
}

// Login godoc
//
//	@Summary		Login User
//	@Description	Sign in for user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.UserLoginRequest	true	"User request is required"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Router			/auth/login [post]
func (uc *UserController) Login(ctx *gin.Context) {
	newUser := model.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(newUser)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	if !valid {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}

	res, err := uc.UserService.Login(newUser)

	if err != nil {
		if err == model.ErrorInvalidEmailOrPassword {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusUnauthorized,
					Message: http.StatusText(http.StatusUnauthorized),
				},
				Error: err.Error(),
			})
			return
		} else if err == model.ErrorInvalidToken {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusInternalServerError,
					Message: http.StatusText(http.StatusInternalServerError),
				},
				Error: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: res,
	})
	return

}
