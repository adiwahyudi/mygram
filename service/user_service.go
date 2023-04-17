package service

import (
	"mygram/helper"
	"mygram/model"
	"mygram/repository"
)

type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (us *UserService) Add(request model.UserRegisterRequest) (model.UserRegisterResponse, error) {

	id := helper.GenerateID()
	hashed_password, err := helper.HashPassword(request.Password)

	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	newUser := model.User{
		ID:       id,
		Email:    request.Email,
		Username: request.Username,
		Password: hashed_password,
		Age:      request.Age,
	}

	res, err := us.UserRepository.Save(newUser)
	if err != nil {
		return model.UserRegisterResponse{}, err
	}

	return model.UserRegisterResponse{
		ID:        res.ID,
		Username:  res.Username,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (us *UserService) Login(request model.UserLoginRequest) (model.UserLoginResponse, error) {

	result, err := us.UserRepository.GetByUsername(request.Username)

	if err != nil {
		return model.UserLoginResponse{}, err
	}

	valid := helper.CheckPasswordHash(request.Password, result.Password)

	if !valid {
		return model.UserLoginResponse{}, model.ErrorInvalidEmailOrPassword
	}

	token, err := helper.GenerateToken(result.ID)
	if err != nil {
		return model.UserLoginResponse{}, model.ErrorInvalidToken
	}

	return model.UserLoginResponse{
		Token: token,
	}, nil
}

func (us *UserService) MyGram(id string) (model.UserGramResponse, error) {
	var photoResponse []model.ListPhotoResponse
	var socialMediaResponse []model.ListSocialMediasResponse
	var commentResponse []model.ListCommentResponse

	result, err := us.UserRepository.GetDetailUser(id)

	if err != nil {
		return model.UserGramResponse{}, err
	}

	for _, val := range result.Photos {
		photoResponse = append(photoResponse, model.ListPhotoResponse{
			ID:        val.ID,
			Title:     val.Title,
			Caption:   val.Caption,
			PhotoURL:  val.PhotoURL,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}
	for _, val := range result.SocialMedias {
		socialMediaResponse = append(socialMediaResponse, model.ListSocialMediasResponse{
			ID:             val.ID,
			Name:           val.Name,
			SocialMediaURL: val.SocialMediaURL,
			CreatedAt:      val.CreatedAt,
			UpdatedAt:      val.UpdatedAt,
		})
	}
	for _, val := range result.Comments {
		commentResponse = append(commentResponse, model.ListCommentResponse{
			ID:        val.ID,
			Message:   val.Message,
			CreatedAt: val.CreatedAt,
			UpdatedAt: val.UpdatedAt,
		})
	}

	return model.UserGramResponse{
		ID:           result.ID,
		Email:        result.Email,
		Username:     result.Username,
		Age:          result.Age,
		Photos:       photoResponse,
		Comments:     commentResponse,
		SocialMedias: socialMediaResponse,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}, nil

}
