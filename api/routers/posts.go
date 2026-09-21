package routers

import (
	"github.com/gin-gonic/gin"
	handler "github.com/vertinofff/blog-api/api/handlers"
)

func Posts(router *gin.RouterGroup, p *handler.PostsHandler, auth gin.HandlerFunc) {
	router.GET("/", p.GetAllPost)
	router.GET("/:id", p.Get)
	protected := router.Group("/", auth)
	protected.POST("/", p.CreatePost)
	protected.PUT("/:id", p.UpdatePost)
	protected.DELETE("/:id", p.DeletePost)
}
