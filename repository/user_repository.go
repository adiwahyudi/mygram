package repository

import (
	"mygram/model"

	"gorm.io/gorm"
)

//go:generate mockery --name IUserRepository
type IUserRepository interface {
	Save(newUser model.User) (model.User, error)
	GetByUsername(username string) (model.User, error)
}
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(newUser model.User) (model.User, error) {
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return newUser, nil
}

func (ur *UserRepository) GetByUsername(username string) (model.User, error) {
	user := model.User{}
	tx := ur.db.First(&user, "username = ?", username)

	return user, tx.Error
}
