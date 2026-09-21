package repositories

import (
	"context"
	"errors"
	"github.com/vertinofff/blog-api/data/models"
	"gorm.io/gorm"
)

type PostRepository struct{ db *gorm.DB }

func NewPostRepository(db *gorm.DB) *PostRepository { return &PostRepository{db: db} }
func (r *PostRepository) List(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	err := r.db.WithContext(ctx).Order("id DESC").Find(&posts).Error
	return posts, err
}
func (r *PostRepository) Get(ctx context.Context, id uint) (*models.Post, error) {
	var p models.Post
	if err := r.db.WithContext(ctx).First(&p, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return &p, nil
}
func (r *PostRepository) Create(ctx context.Context, p *models.Post) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { return tx.Create(p).Error })
}
func (r *PostRepository) Update(ctx context.Context, p *models.Post, version uint) error {
	result := r.db.WithContext(ctx).Model(&models.Post{}).Where("id = ? AND version = ?", p.Id, version).Updates(map[string]interface{}{"title": p.Title, "content": p.Content, "version": version + 1})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	p.Version = version + 1
	return nil
}
func (r *PostRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&models.Post{}, id)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
}
