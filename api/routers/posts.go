package routers

import (
	handler "github.com/vertinofff/blog-api/api/handlers"
	"github.com/vertinofff/blog-api/config"
	"github.com/gin-gonic/gin"
)

func Posts(router *gin.RouterGroup, cfg *config.Config){
	p := handler.NewPostsHandler(cfg)
	router.GET("/" ,p.GetAllPost)
	router.POST("/",p.CreatePost)
	router.GET("/:id" , p.Get)
	router.PUT("/:id",p.UpdatePost)
	router.DELETE("/:id",p.DeletePost)
}