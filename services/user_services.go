package services

import (
	"context"
	"errors"
	"github.com/vertinofff/blog-api/api/dto"
	"github.com/vertinofff/blog-api/data/models"
	"github.com/vertinofff/blog-api/pkg/service_errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	db     *gorm.DB
	tokens *TokenService
}

func NewUserService(db *gorm.DB, tokens *TokenService) *UserService {
	return &UserService{db: db, tokens: tokens}
}
func (s *UserService) LoginByUsername(ctx context.Context, req *dto.LoginByUsernameRequest) (*dto.TokenDetail, error) {
	var u models.User
	err := s.db.WithContext(ctx).Where("username = ?", req.Username).First(&u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, service_errors.ErrInvalidCredentials
		}
		return nil, err
	}
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)) != nil {
		return nil, service_errors.ErrInvalidCredentials
	}
	return s.tokens.GenerateToken(u.Id, u.Username, u.Email)
}
func (s *UserService) RegisterByUsername(ctx context.Context, req *dto.RegisterUserByUsernameRequest) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u := models.User{Username: req.Username, Email: req.Email, Password: string(hash)}
	err = s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error { return tx.Create(&u).Error })
	if err != nil { // Drivers differ; unique constraints are authoritative and mapped here conservatively.
		return &service_errors.ServiceError{Public: "user already exists", Err: service_errors.ErrConflict}
	}
	return nil
}
