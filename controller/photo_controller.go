package controller

import (
	"mygram/model"
	"mygram/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type PhotoController struct {
	PhotoService service.PhotoService
}

func NewPhotoController(photoService service.PhotoService) *PhotoController {
	return &PhotoController{
		PhotoService: photoService,
	}
}

// GetListPhotos godoc
//
//	@Summary		Get All Photo
//	@Description	Get All Photo.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/photo [get]
func (pc *PhotoController) GetListPhotos(ctx *gin.Context) {
	photos, err := pc.PhotoService.GetAll()

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

	ctx.JSON(http.StatusOK, model.ResponseSuccess{
		Meta: model.Meta{
			Code:    http.StatusOK,
			Message: http.StatusText(http.StatusOK),
		},
		Data: photos,
	})
	return
}

// GetPhotoByID godoc
//
//	@Summary		Get Photo by ID.
//	@Description	Get specific Photo by ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/photo/{id} [get]
func (pc *PhotoController) GetPhotoByID(ctx *gin.Context) {
	id := ctx.Param("id")
	photo, err := pc.PhotoService.GetById(id)

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Photo " + err.Error(),
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
		Data: photo,
	})
	return

}

// CreatePhoto godoc
//
//	@Summary		Create Photo
//	@Description	Add new Photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.PhotoCreateRequest	true	"Photo request is required"
//	@Success		201		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/photo [post]
func (pc *PhotoController) CreatePhoto(ctx *gin.Context) {
	newPhoto := model.PhotoCreateRequest{}

	if err := ctx.ShouldBindJSON(&newPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(newPhoto)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	res, err := pc.PhotoService.Add(newPhoto, userId.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, res)
	return

}

// UpdatePhoto godoc
//
//	@Summary		Update Photo
//	@Description	Update photo for specific Photo ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Param			request	body		model.PhotoUpdateRequest	true	"Photo request is required"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		403		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/photo/{id} [put]
func (pc *PhotoController) UpdatePhoto(ctx *gin.Context) {
	updatePhoto := model.PhotoUpdateRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updatePhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(updatePhoto)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: err.Error(),
		})
		return
	}

	updated, err := pc.PhotoService.UpdateById(updatePhoto, id, userId.(string))
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Photo " + err.Error(),
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
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
		Data: updated,
	})
	return
}

// DeletePhoto godoc
//
//	@Summary		Delete Photo
//	@Description	Delete photo for specific Photo ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		403		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/photo/{id} [delete]
func (pc *PhotoController) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			},
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	err := pc.PhotoService.DeleteById(id, userId.(string))

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: "Photo " + err.Error(),
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusForbidden,
					Message: http.StatusText(http.StatusForbidden),
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
		Data: "Delete photo success.",
	})
	return

}
