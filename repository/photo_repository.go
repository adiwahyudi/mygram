package repository

import (
	"errors"
	"mygram/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockery --name IPhotoRepository
type IPhotoRepository interface {
	Get() ([]model.Photo, error)
	GetOne(id string) (model.Photo, error)
	Save(photo model.Photo) (model.Photo, error)
	Update(updatePhoto model.Photo, id string) (model.Photo, error)
	Delete(id string) error
}
type PhotoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *PhotoRepository {
	return &PhotoRepository{
		db: db,
	}
}

func (pr *PhotoRepository) Get() ([]model.Photo, error) {
	photo := make([]model.Photo, 0)

	tx := pr.db.Find(&photo)
	return photo, tx.Error
}

func (pr *PhotoRepository) GetOne(id string) (model.Photo, error) {
	photo := model.Photo{}

	tx := pr.db.First(&photo, "id = ?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return model.Photo{}, model.ErrorNotFound
	}
	return photo, tx.Error
}

func (pr *PhotoRepository) Save(photo model.Photo) (model.Photo, error) {
	tx := pr.db.Create(&photo)
	return photo, tx.Error
}

func (pr *PhotoRepository) Update(updatePhoto model.Photo, id string) (model.Photo, error) {
	tx := pr.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "user_id"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
		},
		).
		Where("id = ?", id).
		Updates(&updatePhoto)
	return updatePhoto, tx.Error
}

func (pr *PhotoRepository) Delete(id string) error {
	photo := model.Photo{
		ID: id,
	}

	tx := pr.db.Select("Comments").Delete(&photo)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
