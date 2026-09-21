package middlewares

import (
	"github.com/gin-gonic/gin"
	helper "github.com/vertinofff/blog-api/api/helpers"
	"github.com/vertinofff/blog-api/constants"
	"github.com/vertinofff/blog-api/services"
	"net/http"
	"strings"
)

func Authentication(tokens *services.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		parts := strings.Fields(c.GetHeader(constants.AuthorizationHeaderKey))
		if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithAnyError(nil, false, helper.AuthError, "unauthorized"))
			return
		}
		claims, e := tokens.GetAccessClaims(parts[1])
		if e != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithAnyError(nil, false, helper.AuthError, "unauthorized"))
			return
		}
		id := uint(claims[constants.UserIdKey].(float64))
		c.Set(constants.UserIdKey, id)
		c.Set(constants.UsernameKey, claims[constants.UsernameKey])
		c.Set(constants.EmailKey, claims[constants.EmailKey])
		c.Next()
	}
}
