package helper

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vertinofff/blog-api/pkg/service_errors"
	"net/http"
)

func WriteError(c *gin.Context, err error) {
	status, code := http.StatusInternalServerError, InternalError
	switch {
	case errors.Is(err, service_errors.ErrValidation):
		status, code = http.StatusBadRequest, ValidationError
	case errors.Is(err, service_errors.ErrInvalidCredentials), errors.Is(err, service_errors.ErrUnauthorized):
		status, code = http.StatusUnauthorized, AuthError
	case errors.Is(err, service_errors.ErrForbidden):
		status, code = http.StatusForbidden, ForbiddenError
	case errors.Is(err, service_errors.ErrNotFound):
		status, code = http.StatusNotFound, NotFoundError
	case errors.Is(err, service_errors.ErrConflict):
		status, code = http.StatusConflict, InternalError
	case errors.Is(err, service_errors.ErrUnavailable):
		status, code = http.StatusServiceUnavailable, InternalError
	}
	c.AbortWithStatusJSON(status, GenerateBaseResponseWithAnyError(nil, false, code, "request could not be completed"))
}
