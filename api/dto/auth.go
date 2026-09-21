package dto

type TokenDetail struct {
	AccessToken            string `json:"accessToken"`
	RefreshToken           string `json:"refreshToken"`
	AccessTokenExpireTime  int64  `json:"accessTokenExpireTime"`
	RefreshTokenExpireTime int64  `json:"refreshTokenExpireTime"`
}

type RegisterUserByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5,max=50"`
	Email    string `json:"email" binding:"required,min=6,max=256,email"`
	Password string `json:"password" binding:"required,password"`
}

type LoginByUsernameRequest struct {
	Username string `json:"username" binding:"required,min=5,max=50"`
	Password string `json:"password" binding:"required,max=64"`
}
