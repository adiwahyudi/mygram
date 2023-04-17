package service

import (
	"mygram/helper"
	"mygram/model"
	"mygram/repository"
)

type PhotoService struct {
	PhotoRepository repository.IPhotoRepository
}

func NewPhotoService(photoRepository repository.IPhotoRepository) *PhotoService {
	return &PhotoService{
		PhotoRepository: photoRepository,
	}
}

func (ps *PhotoService) GetAll() ([]model.PhotoResponseBaru, error) {
	photosResponse := make([]model.PhotoResponseBaru, 0)

	res, err := ps.PhotoRepository.Get()

	if err != nil {
		return []model.PhotoResponseBaru{}, err
	}

	for _, val := range res {
		commentResponse := make([]model.CommentInPhotoResponse, 0)
		for _, comment := range val.Comments {
			commentResponse = append(commentResponse, model.CommentInPhotoResponse{
				ID:        comment.ID,
				UserID:    comment.UserID,
				Message:   comment.Message,
				CreatedAt: comment.CreatedAt,
				UpdatedAt: comment.UpdatedAt,
			})
		}
		photosResponse = append(photosResponse, model.PhotoResponseBaru{
			ID:        val.ID,
			UserID:    val.UserID,
			Title:     val.Title,
			Caption:   val.Caption,
			PhotoURL:  val.PhotoURL,
			Comments:  commentResponse,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return photosResponse, nil
}

func (ps *PhotoService) GetById(id string) (model.PhotoResponse, error) {
	photo, err := ps.PhotoRepository.GetOne(id)

	if err != nil {
		return model.PhotoResponse{}, err
	}

	return model.PhotoResponse{
		ID:        photo.ID,
		UserID:    photo.UserID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		CreatedAt: photo.CreatedAt,
		UpdatedAt: photo.UpdatedAt,
	}, nil
}

func (ps *PhotoService) Add(request model.PhotoCreateRequest, userId string) (model.PhotoCreateResponse, error) {
	id := helper.GenerateID()
	photo := model.Photo{
		ID:       id,
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
		UserID:   userId,
	}

	res, err := ps.PhotoRepository.Save(photo)
	if err != nil {
		return model.PhotoCreateResponse{}, err
	}

	return model.PhotoCreateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoURL:  res.PhotoURL,
		CreatedAt: res.CreatedAt,
	}, nil

}

func (ps *PhotoService) UpdateById(request model.PhotoUpdateRequest, id string, userId string) (model.PhotoUpdateResponse, error) {
	getById, err := ps.PhotoRepository.GetOne(id)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.PhotoUpdateResponse{}, err
		}
		return model.PhotoUpdateResponse{}, model.ErrorNotFound
	}

	if getById.UserID != userId {
		return model.PhotoUpdateResponse{}, model.ErrorForbiddenAccess
	}

	photo := model.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoURL: request.PhotoURL,
	}

	res, err := ps.PhotoRepository.Update(photo, id)
	if err != nil {
		return model.PhotoUpdateResponse{}, err
	}

	return model.PhotoUpdateResponse{
		ID:        res.ID,
		UserID:    res.UserID,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoURL:  res.PhotoURL,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (ps *PhotoService) DeleteById(id string, userId string) error {
	getById, err := ps.PhotoRepository.GetOne(id)
	if err != nil {
		if err != model.ErrorNotFound {
			return err
		}
		return model.ErrorNotFound
	}

	if getById.UserID != userId {
		return model.ErrorForbiddenAccess
	}

	err = ps.PhotoRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
