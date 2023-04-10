package controller

import (
	"mygram/model"
	"mygram/service"
	"net/http"

	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) *CommentController {
	return &CommentController{
		CommentService: commentService,
	}
}

// GetListComments godoc
//
//	@Summary		Get all comment
//	@Description	View all comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Success		200		{array}		model.CommentResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/comment [get]
func (cc *CommentController) GetListComments(ctx *gin.Context) {
	comments, err := cc.CommentService.GetAll()

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comments)
	return
}

// GetOneCommentsByID godoc
//
//	@Summary		Get comment by ID
//	@Description	View specific comment by ID
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Comment ID"
//	@Success		200		{object}	model.CommentResponse
//	@Failure		401		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/comment/:id [get]
func (cc *CommentController) GetOneCommentsByID(ctx *gin.Context) {
	id := ctx.Param("id")

	comment, err := cc.CommentService.GetById(id)
	if err != nil {
		if err != model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
				Err: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
			Err: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
	return
}

// CreateCommentByPhotoID godoc
//
//	@Summary		Create comment
//	@Description	Add new comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			photo_id	path		string	true	"Photo ID"
//	@Param			request	body		model.CommentCreateRequest	true	"Comment request is required"
//	@Success		201		{object}	model.CommentCreateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/comment/:photo_id [post]
func (cc *CommentController) CreateCommentByPhotoID(ctx *gin.Context) {
	commentRequest := model.CommentCreateRequest{}

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(commentRequest)
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

	photoId := ctx.Param("photo_id")
	result, err := cc.CommentService.Add(commentRequest, userId.(string), photoId)

	if err != nil {
		if err != model.ErrorNotFound {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
				Err: err.Error(),
			})
			return
		}
		ctx.AbortWithStatusJSON(http.StatusNotFound, model.MyError{
			Err: model.ErrorPhotoNotFound.Err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, result)
	return

}

// UpdateComment godoc
//
//	@Summary		Update comment
//	@Description	Update specific comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Comment ID"
//	@Param			request	body		model.CommentUpdateRequest	true	"Comment request is required"
//	@Success		200		{object}	model.CommentUpdateResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/comment/:id [put]
func (cc *CommentController) UpdateComment(ctx *gin.Context) {
	commentRequest := model.CommentUpdateRequest{}

	if err := ctx.ShouldBindJSON(&commentRequest); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, model.MyError{
			Err: err.Error(),
		})
		return
	}

	valid, err := valid.ValidateStruct(commentRequest)
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

	id := ctx.Param("id")
	result, err := cc.CommentService.UpdateById(commentRequest, userId.(string), id)

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

	ctx.JSON(http.StatusOK, result)
	return
}

// DeleteComment godoc
//
//	@Summary		Delete comment
//	@Description	Delete comment
//	@Tags			Comment
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Comment ID"
//	@Success		200		{object}	model.DeleteCommentResponse
//	@Failure		400		{object}	model.MyError
//	@Failure		401		{object}	model.MyError
//	@Failure		403		{object}	model.MyError
//	@Failure		404		{object}	model.MyError
//	@Failure		500		{object}	model.MyError
//	@Security		Bearer
//	@Router			/comment/:id [delete]
func (cc *CommentController) DeleteComment(ctx *gin.Context) {

	userId, isExist := ctx.Get("user_id")
	if !isExist {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, model.MyError{
			Err: model.ErrorInvalidToken.Err,
		})
		return
	}

	id := ctx.Param("id")
	err := cc.CommentService.DeleteById(userId.(string), id)

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

	ctx.JSON(http.StatusOK, model.DeleteCommentResponse{
		Message: "Delete comment success!",
	})
	return
}
