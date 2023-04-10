package service

import (
	"mygram/helper"
	"mygram/model"
	"mygram/repository"
)

type SocialMediaService struct {
	SocialMediaRepository repository.ISocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.ISocialMediaRepository) *SocialMediaService {
	return &SocialMediaService{
		SocialMediaRepository: socialMediaRepository,
	}
}

func (sms *SocialMediaService) GetAll() ([]model.SocialMediaResponse, error) {
	socialMediaRespons := make([]model.SocialMediaResponse, 0)

	res, err := sms.SocialMediaRepository.Get()

	if err != nil {
		return []model.SocialMediaResponse{}, err
	}

	for _, val := range res {
		socialMediaRespons = append(socialMediaRespons, model.SocialMediaResponse{
			ID:             val.ID,
			UserID:         val.UserID,
			Name:           val.Name,
			SocialMediaURL: val.SocialMediaURL,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}

	return socialMediaRespons, nil
}

func (sms *SocialMediaService) GetById(id string) (model.SocialMediaResponse, error) {
	socialMedia, err := sms.SocialMediaRepository.GetOne(id)

	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaResponse{}, err
		}
		return model.SocialMediaResponse{}, model.ErrorNotFound
	}

	return model.SocialMediaResponse{
		ID:             socialMedia.ID,
		UserID:         socialMedia.UserID,
		Name:           socialMedia.Name,
		SocialMediaURL: socialMedia.SocialMediaURL,
		CreatedAt:      socialMedia.CreatedAt,
		UpdatedAt:      socialMedia.UpdatedAt,
	}, nil
}

func (sms *SocialMediaService) Add(request model.SocialMediaCreateRequest, userId string) (model.SocialMediaCreateResponse, error) {
	id := helper.GenerateID()
	socialMedia := model.SocialMedia{
		ID:             id,
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
		UserID:         userId,
	}

	res, err := sms.SocialMediaRepository.Save(socialMedia)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaCreateResponse{}, model.ErrorNotFound
		}
		return model.SocialMediaCreateResponse{}, err
	}

	return model.SocialMediaCreateResponse{
		ID:             res.ID,
		UserID:         res.UserID,
		Name:           res.Name,
		SocialMediaURL: res.SocialMediaURL,
		CreatedAt:      res.CreatedAt,
	}, nil

}

func (sms *SocialMediaService) UpdateById(request model.SocialMediaUpdateRequest, id string, userId string) (model.SocialMediaUpdateResponse, error) {
	getById, err := sms.SocialMediaRepository.GetOne(id)
	if err != nil {
		if err != model.ErrorNotFound {
			return model.SocialMediaUpdateResponse{}, err
		}
		return model.SocialMediaUpdateResponse{}, model.ErrorNotFound
	}

	if getById.UserID != userId {
		return model.SocialMediaUpdateResponse{}, model.ErrorForbiddenAccess
	}

	socialMedia := model.SocialMedia{
		Name:           request.Name,
		SocialMediaURL: request.SocialMediaURL,
	}

	res, err := sms.SocialMediaRepository.Update(socialMedia, id)
	if err != nil {
		return model.SocialMediaUpdateResponse{}, err
	}

	return model.SocialMediaUpdateResponse{
		ID:             res.ID,
		UserID:         res.UserID,
		Name:           res.Name,
		SocialMediaURL: res.SocialMediaURL,

		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (sms *SocialMediaService) DeleteById(id string, userId string) error {
	getById, err := sms.SocialMediaRepository.GetOne(id)
	if err != nil {
		if err != model.ErrorNotFound {
			return err
		}
		return model.ErrorNotFound
	}

	if getById.UserID != userId {
		return model.ErrorForbiddenAccess
	}

	err = sms.SocialMediaRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
