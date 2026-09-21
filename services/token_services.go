package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/vertinofff/blog-api/api/dto"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/constants"
	"github.com/vertinofff/blog-api/pkg/service_errors"
	"time"
)

type TokenService struct{ cfg *config.Config }

func NewTokenService(cfg *config.Config) *TokenService { return &TokenService{cfg: cfg} }
func (s *TokenService) GenerateToken(userID int, username, email string) (*dto.TokenDetail, error) {
	now := time.Now()
	td := &dto.TokenDetail{AccessTokenExpireTime: now.Add(s.cfg.JWT.AccessTokenExpireDuration * time.Minute).Unix(), RefreshTokenExpireTime: now.Add(s.cfg.JWT.RefreshTokenExpireDuration * time.Minute).Unix()}
	access := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{constants.UserIdKey: userID, constants.UsernameKey: username, constants.EmailKey: email, constants.ExpireTimeKey: td.AccessTokenExpireTime, "typ": "access"})
	var err error
	td.AccessToken, err = access.SignedString([]byte(s.cfg.JWT.Secret))
	if err != nil {
		return nil, err
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{constants.UserIdKey: userID, constants.ExpireTimeKey: td.RefreshTokenExpireTime, "typ": "refresh"})
	td.RefreshToken, err = refresh.SignedString([]byte(s.cfg.JWT.RefreshSecret))
	if err != nil {
		return nil, err
	}
	return td, nil
}
func (s *TokenService) GetAccessClaims(raw string) (map[string]interface{}, error) {
	token, err := jwt.Parse(raw, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, service_errors.ErrUnauthorized
		}
		return []byte(s.cfg.JWT.Secret), nil
	})
	if err != nil {
		return nil, service_errors.ErrUnauthorized
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid || claims["typ"] != "access" {
		return nil, service_errors.ErrUnauthorized
	}
	if _, ok := claims[constants.UserIdKey].(float64); !ok {
		return nil, service_errors.ErrUnauthorized
	}
	return claims, nil
}
func IsExpired(err error) bool {
	var v *jwt.ValidationError
	return errors.As(err, &v) && v.Errors&jwt.ValidationErrorExpired != 0
}
