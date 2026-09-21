package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	handler "github.com/vertinofff/blog-api/api/handlers"
	"github.com/vertinofff/blog-api/api/middlewares"
	"github.com/vertinofff/blog-api/api/routers"
	validation "github.com/vertinofff/blog-api/api/validations"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/data/repositories"
	"github.com/vertinofff/blog-api/services"
	"gorm.io/gorm"
)

func NewServer(cfg *config.Config, database *gorm.DB) *http.Server {
	tokens := services.NewTokenService(cfg)
	users := services.NewUserService(database, tokens)
	posts := services.NewPostService(repositories.NewPostRepository(database))
	r := gin.New()
	r.Use(gin.Recovery(), maxBody(1<<20))
	RegisterValidators(cfg)
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })
	r.GET("/readyz", func(c *gin.Context) {
		if e := database.WithContext(c.Request.Context()).Exec("SELECT 1").Error; e != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		c.Status(http.StatusOK)
	})
	v1 := r.Group("/api/v1")
	routers.Auth(v1.Group("/auth"), handler.NewUsersHandler(users))
	routers.Posts(v1.Group("/posts"), handler.NewPostsHandler(posts), middlewares.Authentication(tokens))
	return &http.Server{Addr: ":" + cfg.Server.InternalPort, Handler: r, ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second, IdleTimeout: 60 * time.Second}
}
func Run(ctx context.Context, s *http.Server) error {
	go func() {
		<-ctx.Done()
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = s.Shutdown(shutdownCtx)
	}()
	if e := s.ListenAndServe(); e != nil && !errors.Is(e, http.ErrServerClosed) {
		return fmt.Errorf("serve: %w", e)
	}
	return nil
}
func maxBody(n int64) gin.HandlerFunc {
	return func(c *gin.Context) { c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, n); c.Next() }
}
func RegisterValidators(cfg *config.Config) {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if e := val.RegisterValidation("password", validation.PasswordValidator(cfg.Password), true); e != nil {
			panic(fmt.Errorf("register password validation: %w", e))
		}
	}
}
