package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vertinofff/blog-api/api/dto"
	helper "github.com/vertinofff/blog-api/api/helpers"
	"github.com/vertinofff/blog-api/services"
	"net/http"
)

type AuthHandler struct{ service *services.UserService }

func NewUsersHandler(service *services.UserService) *AuthHandler {
	return &AuthHandler{service: service}
}
func (h *AuthHandler) Login(c *gin.Context) {
	req := new(dto.LoginByUsernameRequest)
	if e := c.ShouldBindJSON(req); e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, e))
		return
	}
	token, e := h.service.LoginByUsername(c.Request.Context(), req)
	if e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, helper.Success))
}
func (h *AuthHandler) Register(c *gin.Context) {
	req := new(dto.RegisterUserByUsernameRequest)
	if e := c.ShouldBindJSON(req); e != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, e))
		return
	}
	if e := h.service.RegisterByUsername(c.Request.Context(), req); e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, helper.Success))
}
