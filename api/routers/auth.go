package routers

import (
	"github.com/gin-gonic/gin"
	handler "github.com/vertinofff/blog-api/api/handlers"
)

func Auth(router *gin.RouterGroup, h *handler.AuthHandler) {
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
}
