package routers

import (
	handler "github.com/vertinofff/blog-api/api/handlers"
	"github.com/vertinofff/blog-api/config"
	"github.com/gin-gonic/gin"
)

func Auth(router *gin.RouterGroup, cfg *config.Config) {
	h := handler.NewUsersHandler(cfg)
	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
}