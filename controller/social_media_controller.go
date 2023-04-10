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

// GetList godoc
//
//	@Summary		Get All Social Media
//	@Description	Get All Social Media.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.SocialMediaResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/social_media [get]
func (smc *SocialMediaController) GetList(ctx *gin.Context) {
	socialMedias, err := smc.SocialMediaService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedias)
	return
}

// GetByID godoc
//
//	@Summary		Get Social Media by ID.
//	@Description	Get specific Social Media by ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Success		200		{object}	model.SocialMediaResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/social_media/{id} [get]
func (smc *SocialMediaController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := smc.SocialMediaService.GetById(id)

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
//	@Summary		Create Social Media
//	@Description	Add new Social Media
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.SocialMediaCreateRequest	true	"Social Media request is required"
//	@Success		201		{object}	model.SocialMediaCreateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/social_media [post]
func (smc *SocialMediaController) Create(ctx *gin.Context) {
	newSocialMedia := model.SocialMediaCreateRequest{}

	if err := ctx.ShouldBindJSON(&newSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(newSocialMedia)

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

	res, err := smc.SocialMediaService.Add(newSocialMedia, userId.(string))
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
//	@Summary		Update Social Media
//	@Description	Update Social Media for specific Social Media ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Param			request	body		model.SocialMediaUpdateRequest	true	"Social Media request is required"
//	@Success		200		{object}	model.SocialMediaUpdateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError``
//	@Security		Bearer
//	@Router			/social_media/{id} [put]
func (smc *SocialMediaController) Update(ctx *gin.Context) {
	updateSocialMedia := model.SocialMediaUpdateRequest{}
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&updateSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}
	valid, err := valid.ValidateStruct(updateSocialMedia)

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

	res, err := smc.SocialMediaService.UpdateById(updateSocialMedia, id, userId.(string))
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
//	@Summary		Delete Social Media
//	@Description	Delete Social Media for specific Social Media ID.
//	@Tags			Social Media
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Social Media ID"
//	@Success		200		{object}	model.DeleteSocialMediaResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/social_media/{id} [delete]
func (smc *SocialMediaController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	err := smc.SocialMediaService.DeleteById(id, userId.(string))

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

	ctx.JSON(http.StatusOK, model.DeleteSocialMediaResponse{
		Message: "Success delete Social Media!",
	})

}
