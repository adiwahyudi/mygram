package repository

import (
	"errors"
	"mygram/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate mockery --name ICommentRepository
type ICommentRepository interface {
	Get() ([]model.Comment, error)
	GetOne(id string) (model.Comment, error)
	Save(comment model.Comment) (model.Comment, error)
	Update(updateComment model.Comment, id string) (model.Comment, error)
	Delete(id string) error
}
type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) *CommentRepository {
	return &CommentRepository{
		db: db,
	}
}

func (cr *CommentRepository) Get() ([]model.Comment, error) {
	comment := make([]model.Comment, 0)

	tx := cr.db.Find(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) GetOne(id string) (model.Comment, error) {
	comment := model.Comment{}

	tx := cr.db.First(&comment, "id = ?", id)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return model.Comment{}, model.ErrorNotFound
	}
	return comment, tx.Error
}

func (cr *CommentRepository) Save(comment model.Comment) (model.Comment, error) {
	tx := cr.db.Create(&comment)
	return comment, tx.Error
}

func (cr *CommentRepository) Update(updateComment model.Comment, id string) (model.Comment, error) {
	tx := cr.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "user_id"},
				{Name: "photo_id"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
		},
		).
		Where("id = ?", id).
		Updates(&updateComment)
	return updateComment, tx.Error
}

func (cr *CommentRepository) Delete(id string) error {
	comment := model.Comment{}

	tx := cr.db.Delete(&comment, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
