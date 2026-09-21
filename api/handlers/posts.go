package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vertinofff/blog-api/api/dto"
	helper "github.com/vertinofff/blog-api/api/helpers"
	"github.com/vertinofff/blog-api/constants"
	"github.com/vertinofff/blog-api/pkg/service_errors"
	"github.com/vertinofff/blog-api/services"
	"net/http"
	"strconv"
)

type PostsHandler struct{ service *services.PostService }

func NewPostsHandler(s *services.PostService) *PostsHandler { return &PostsHandler{service: s} }
func postID(c *gin.Context) (uint, bool) {
	id, e := strconv.ParseUint(c.Param("id"), 10, 32)
	if e != nil || id == 0 {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponse(nil, false, helper.ValidationError))
		return 0, false
	}
	return uint(id), true
}
func userID(c *gin.Context) (uint, bool) {
	v, ok := c.Get(constants.UserIdKey)
	if !ok {
		return 0, false
	}
	id, ok := v.(uint)
	return id, ok
}
func (p *PostsHandler) GetAllPost(c *gin.Context) {
	posts, e := p.service.List(c.Request.Context())
	if e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(posts, true, helper.Success))
}
func (p *PostsHandler) CreatePost(c *gin.Context) {
	req := new(dto.CreatePost)
	if e := c.ShouldBindJSON(req); e != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, e))
		return
	}
	author, ok := userID(c)
	if !ok {
		helper.WriteError(c, service_errors.ErrUnauthorized)
		return
	}
	post, e := p.service.Create(c.Request.Context(), author, req)
	if e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(post, true, helper.Success))
}
func (p *PostsHandler) Get(c *gin.Context) {
	id, ok := postID(c)
	if !ok {
		return
	}
	post, e := p.service.Get(c.Request.Context(), id)
	if e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(post, true, helper.Success))
}
func (p *PostsHandler) UpdatePost(c *gin.Context) {
	id, ok := postID(c)
	if !ok {
		return
	}
	req := new(dto.UpdatePost)
	if e := c.ShouldBindJSON(req); e != nil {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, e))
		return
	}
	author, ok := userID(c)
	if !ok {
		helper.WriteError(c, service_errors.ErrUnauthorized)
		return
	}
	post, e := p.service.Update(c.Request.Context(), id, author, req)
	if e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(post, true, helper.Success))
}
func (p *PostsHandler) DeletePost(c *gin.Context) {
	id, ok := postID(c)
	if !ok {
		return
	}
	author, ok := userID(c)
	if !ok {
		helper.WriteError(c, service_errors.ErrUnauthorized)
		return
	}
	if e := p.service.Delete(c.Request.Context(), id, author); e != nil {
		helper.WriteError(c, e)
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("deleted", true, helper.Success))
}
