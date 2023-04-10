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

// GetList godoc
//
//	@Summary		Get All Photo
//	@Description	Get All Photo.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.PhotoResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/photo [get]
func (pc *PhotoController) GetList(ctx *gin.Context) {
	photos, err := pc.PhotoService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
	return
}

// GetByID godoc
//
//	@Summary		Get Photo by ID.
//	@Description	Get specific Photo by ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Success		200		{object}	model.PhotoResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/photo/{id} [get]
func (pc *PhotoController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := pc.PhotoService.GetById(id)

	if err != nil {
		if err != model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
				Err: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
			Err: model.ErrorNotFound.Err,
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
	return

}

// Create godoc
//
//	@Summary		Create Photo
//	@Description	Add new Photo
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.PhotoCreateRequest	true	"Photo request is required"
//	@Success		201		{object}	model.PhotoCreateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/photo [post]
func (pc *PhotoController) Create(ctx *gin.Context) {
	newPhoto := model.PhotoCreateRequest{}

	if err := ctx.ShouldBindJSON(&newPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(newPhoto)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
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

// Update godoc
//
//	@Summary		Update Photo
//	@Description	Update photo for specific Photo ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Param			request	body		model.PhotoUpdateRequest	true	"Photo request is required"
//	@Success		200		{object}	model.PhotoUpdateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError``
//	@Security		Bearer
//	@Router			/photo/{id} [put]
func (pc *PhotoController) Update(ctx *gin.Context) {
	updatePhoto := model.PhotoUpdateRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updatePhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(updatePhoto)

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

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := pc.PhotoService.UpdateById(updatePhoto, id, userId.(string))
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
				Err: model.ErrorNotFound.Err,
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.MyError{
				Err: model.ErrorForbiddenAccess.Err,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, res)
	return
}

// Delete godoc
//
//	@Summary		Delete Photo
//	@Description	Delete photo for specific Photo ID.
//	@Tags			Photo
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Photo ID"
//	@Success		200		{object}	model.DeletePhotoResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/photo/{id} [delete]
func (pc *PhotoController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	err := pc.PhotoService.DeleteById(id, userId.(string))

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
				Err: model.ErrorNotFound.Err,
			})
			return
		} else if err == model.ErrorForbiddenAccess {
			ctx.AbortWithStatusJSON(http.StatusForbidden, model.MyError{
				Err: model.ErrorForbiddenAccess.Err,
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, model.DeletePhotoResponse{
		Message: "Success delete Photo!",
	})

}
