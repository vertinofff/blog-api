package services

import (
	"context"
	"errors"
	"github.com/vertinofff/blog-api/api/dto"
	"github.com/vertinofff/blog-api/data/models"
	"github.com/vertinofff/blog-api/data/repositories"
	"github.com/vertinofff/blog-api/pkg/service_errors"
	"gorm.io/gorm"
)

type PostService struct{ repo *repositories.PostRepository }

func NewPostService(repo *repositories.PostRepository) *PostService    { return &PostService{repo: repo} }
func (s *PostService) List(ctx context.Context) ([]models.Post, error) { return s.repo.List(ctx) }
func (s *PostService) Get(ctx context.Context, id uint) (*models.Post, error) {
	p, e := s.repo.Get(ctx, id)
	if errors.Is(e, gorm.ErrRecordNotFound) {
		return nil, service_errors.ErrNotFound
	}
	return p, e
}
func (s *PostService) Create(ctx context.Context, authorID uint, req *dto.CreatePost) (*models.Post, error) {
	p := &models.Post{Title: req.Title, Content: req.Content, AuthorId: authorID}
	if e := s.repo.Create(ctx, p); e != nil {
		return nil, e
	}
	return p, nil
}
func (s *PostService) Update(ctx context.Context, id, authorID uint, req *dto.UpdatePost) (*models.Post, error) {
	p, e := s.Get(ctx, id)
	if e != nil {
		return nil, e
	}
	if p.AuthorId != authorID {
		return nil, service_errors.ErrForbidden
	}
	p.Title = req.Title
	p.Content = req.Content
	if e = s.repo.Update(ctx, p, p.Version); e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, service_errors.ErrConflict
		}
		return nil, e
	}
	return p, nil
}
func (s *PostService) Delete(ctx context.Context, id, authorID uint) error {
	p, e := s.Get(ctx, id)
	if e != nil {
		return e
	}
	if p.AuthorId != authorID {
		return service_errors.ErrForbidden
	}
	if e = s.repo.Delete(ctx, id); errors.Is(e, gorm.ErrRecordNotFound) {
		return service_errors.ErrNotFound
	}
	return e
}
