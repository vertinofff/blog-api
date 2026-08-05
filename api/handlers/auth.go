package handler

import (
	"net/http"

	"github.com/vertinofff/blog-api/api/dto"
	helper "github.com/vertinofff/blog-api/api/helpers"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/services"
	"github.com/gin-gonic/gin"
	
)

type AuthHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *AuthHandler {
	service := services.NewUserService(cfg)
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	token, err := h.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}

func (h *AuthHandler) Register(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}
