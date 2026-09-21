package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/vertinofff/blog-api/common"
	"github.com/vertinofff/blog-api/config"
)

func PasswordValidator(policy config.PasswordConfig) validator.Func {
	return func(fld validator.FieldLevel) bool {
		value, ok := fld.Field().Interface().(string)
		return ok && common.CheckPassword(value, policy)
	}
}
