package service

import (
	"fmt"
	"mygram/helper"
	"mygram/model"
	"mygram/repository"
)

type CommentService struct {
	CommentRepository repository.ICommentRepository
	PhotoRepository   repository.IPhotoRepository
}

func NewCommentService(commentRepository repository.ICommentRepository, photoRepository repository.IPhotoRepository) *CommentService {
	return &CommentService{
		CommentRepository: commentRepository,
		PhotoRepository:   photoRepository,
	}
}

func (cs *CommentService) Add(request model.CommentCreateRequest, userId string, photoId string) (model.CommentCreateResponse, error) {
	id := helper.GenerateID()

	fmt.Println(photoId)
	_, err := cs.PhotoRepository.GetOne(photoId)
	if err != nil {
		return model.CommentCreateResponse{}, err
	}

	comment := model.Comment{
		ID:      id,
		UserID:  userId,
		PhotoID: photoId,
		Message: request.Message,
	}

	res, err := cs.CommentRepository.Save(comment)
	if err != nil {
		return model.CommentCreateResponse{}, err
	}

	return model.CommentCreateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		Message:   res.Message,
		CreatedAt: res.CreatedAt,
	}, nil
}

func (cs *CommentService) GetAll() ([]model.CommentResponse, error) {
	commentResponse := make([]model.CommentResponse, 0)

	res, err := cs.CommentRepository.Get()

	if err != nil {
		return []model.CommentResponse{}, err
	}

	for _, val := range res {
		commentResponse = append(commentResponse, model.CommentResponse{
			ID:        val.ID,
			UserID:    val.UserID,
			PhotoID:   val.PhotoID,
			Message:   val.Message,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return commentResponse, nil
}

func (cs *CommentService) GetById(id string) (model.CommentResponse, error) {
	res, err := cs.CommentRepository.GetOne(id)

	if err != nil {
		if err == model.ErrorNotFound {
			return model.CommentResponse{}, model.ErrorNotFound
		}
		return model.CommentResponse{}, err
	}

	return model.CommentResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		Message:   res.Message,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (cs *CommentService) UpdateById(request model.CommentUpdateRequest, userId string, id string) (model.CommentUpdateResponse, error) {
	comment, err := cs.CommentRepository.GetOne(id)

	if err != nil {
		if err == model.ErrorNotFound {
			return model.CommentUpdateResponse{}, model.ErrorNotFound
		}
		return model.CommentUpdateResponse{}, err
	}

	if comment.UserID != userId {
		return model.CommentUpdateResponse{}, model.ErrorForbiddenAccess
	}

	commentUpdate := model.Comment{
		Message: request.Message,
	}

	res, err := cs.CommentRepository.Update(commentUpdate, id)
	if err != nil {
		return model.CommentUpdateResponse{}, err
	}

	return model.CommentUpdateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		PhotoID:   res.PhotoID,
		Message:   res.Message,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil

}

func (cs *CommentService) DeleteById(userId string, id string) error {

	comment, err := cs.CommentRepository.GetOne(id)

	if err != nil {
		if err == model.ErrorNotFound {
			return model.ErrorNotFound
		}
		return err
	}

	if comment.UserID != userId {
		return model.ErrorForbiddenAccess
	}

	err = cs.CommentRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
