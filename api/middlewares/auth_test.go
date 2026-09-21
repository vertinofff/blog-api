package middlewares

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/vertinofff/blog-api/config"
	"github.com/vertinofff/blog-api/services"
)

func TestAuthenticationRejectsMalformedHeader(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tokens := services.NewTokenService(&config.Config{JWT: config.JWTConfig{Secret: "access-secret-which-is-long-enough", RefreshSecret: "refresh-secret-which-is-long-enough"}})
	r := gin.New()
	r.GET("/", Authentication(tokens), func(c *gin.Context) { c.Status(http.StatusNoContent) })
	for _, header := range []string{"", "Bearer", "Basic abc", "Bearer a b"} {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Authorization", header)
		res := httptest.NewRecorder()
		r.ServeHTTP(res, req)
		if res.Code != http.StatusUnauthorized {
			t.Fatalf("header %q: got %d", header, res.Code)
		}
	}
}
