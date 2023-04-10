package controller

import (
	"mygram/model"
	"mygram/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type SocialMediaController struct {
	SocialMediaService service.SocialMediaService
}

func NewSocialMediaController(socialMediaService service.SocialMediaService) *SocialMediaController {
	return &SocialMediaController{
		SocialMediaService: socialMediaService,
	}
}

// GetListSocialMedias godoc
//
//	@Summary		Get All Social Media
//	@Description	Get All Social Media.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}		model.ResponseSuccess
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/social_media [get]
func (smc *SocialMediaController) GetListSocialMedias(ctx *gin.Context) {
	socialMedias, err := smc.SocialMediaService.GetAll()
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
		Data: socialMedias,
	})
	return
}

// GetOneSocialMediaByID godoc
//
//	@Summary		Get Social Media by ID.
//	@Description	Get specific Social Media by ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/social_media/{id} [get]
func (smc *SocialMediaController) GetOneSocialMediaByID(ctx *gin.Context) {
	id := ctx.Param("id")
	socialMedia, err := smc.SocialMediaService.GetById(id)

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
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
		Data: socialMedia,
	})
	return
}

// CreateSocialMedia godoc
//
//	@Summary		Create Social Media
//	@Description	Add new Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.SocialMediaCreateRequest	true	"Social Media request is required"
//	@Success		201		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/social_media [post]
func (smc *SocialMediaController) CreateSocialMedia(ctx *gin.Context) {
	newSocialMedia := model.SocialMediaCreateRequest{}

	if err := ctx.ShouldBindJSON(&newSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(newSocialMedia)

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
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := smc.SocialMediaService.Add(newSocialMedia, userId.(string))
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

// UpdateSocialMedia godoc
//
//	@Summary		Update Social Media
//	@Description	Update Social Media for specific Social Media ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Param			request	body		model.SocialMediaUpdateRequest	true	"Social Media request is required"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		403		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/social_media/{id} [put]
func (smc *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	updateSocialMedia := model.SocialMediaUpdateRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updateSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.ResponseFailed{
			Meta: model.Meta{
				Code:    http.StatusBadRequest,
				Message: http.StatusText(http.StatusBadRequest),
			},
			Error: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(updateSocialMedia)

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
			Error: model.ErrorInvalidToken.Err,
		})
		return
	}

	res, err := smc.SocialMediaService.UpdateById(updateSocialMedia, id, userId.(string))
	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
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
		Data: res,
	})
	return
}

// DeleteSocialMedia godoc
//
//	@Summary		Delete Social Media
//	@Description	Delete Social Media for specific Social Media ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Success		200		{object}	model.ResponseSuccess
//	@Failure		400		{object}	model.ResponseFailed
//	@Failure		401		{object}	model.ResponseFailed
//	@Failure		403		{object}	model.ResponseFailed
//	@Failure		404		{object}	model.ResponseFailed
//	@Failure		500		{object}	model.ResponseFailed
//	@Security		Bearer
//	@Router			/social_media/{id} [delete]
func (smc *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
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

	err := smc.SocialMediaService.DeleteById(id, userId.(string))

	if err != nil {
		if err == model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusNotFound, model.ResponseFailed{
				Meta: model.Meta{
					Code:    http.StatusNotFound,
					Message: http.StatusText(http.StatusNotFound),
				},
				Error: err.Error(),
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
		Data: "Delete social media success.",
	})
	return

}
