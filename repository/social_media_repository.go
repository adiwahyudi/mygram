package repository

import (
	"errors"
	"mygram/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockery --name ISocialMediaRepository
type ISocialMediaRepository interface {
	Get() ([]model.SocialMedia, error)
	GetOne(id string) (model.SocialMedia, error)
	Save(socialMedia model.SocialMedia) (model.SocialMedia, error)
	Update(updateSocialMedia model.SocialMedia, id string) (model.SocialMedia, error)
	Delete(id string) error
}
type SocialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) *SocialMediaRepository {
	return &SocialMediaRepository{
		db: db,
	}
}

func (smr *SocialMediaRepository) Get() ([]model.SocialMedia, error) {
	socialMedia := make([]model.SocialMedia, 0)

	// tx := smr.db.Model(&model.SocialMedia{}).Preload("CreditCards").Find(&socialMedia)
	tx := smr.db.Find(&socialMedia)
	return socialMedia, tx.Error
}

func (smr *SocialMediaRepository) GetOne(id string) (model.SocialMedia, error) {
	socialMedia := model.SocialMedia{}

	tx := smr.db.First(&socialMedia, "id = ?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return model.SocialMedia{}, model.ErrorNotFound
	}
	return socialMedia, tx.Error
}

func (smr *SocialMediaRepository) Save(socialMedia model.SocialMedia) (model.SocialMedia, error) {
	tx := smr.db.Create(&socialMedia)
	return socialMedia, tx.Error
}

func (smr *SocialMediaRepository) Update(updateSocialMedia model.SocialMedia, id string) (model.SocialMedia, error) {
	tx := smr.db.
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
		Updates(&updateSocialMedia)
	return updateSocialMedia, tx.Error
}

func (smr *SocialMediaRepository) Delete(id string) error {
	socialMedia := model.SocialMedia{}

	tx := smr.db.Delete(&socialMedia, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
