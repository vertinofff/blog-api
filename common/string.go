package common

import (
	"crypto/rand"
	"github.com/vertinofff/blog-api/config"
	"math/big"
	"regexp"
	"strings"
	"unicode"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	matchFirstCap  = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap    = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func CheckPassword(password string, policy config.PasswordConfig) bool {
	if len(password) < policy.MinLength || len(password) > policy.MaxLength {
		return false
	}
	if policy.IncludeChars && !HasLetter(password) {
		return false
	}
	if policy.IncludeDigits && !HasDigits(password) {
		return false
	}
	if policy.IncludeLowercase && !HasLower(password) {
		return false
	}
	return !policy.IncludeUppercase || HasUpper(password)
}
func GeneratePassword(length int) (string, error) {
	if length < 12 {
		length = 12
	}
	var b strings.Builder
	for i := 0; i < length; i++ {
		n, e := rand.Int(rand.Reader, big.NewInt(int64(len(allCharSet))))
		if e != nil {
			return "", e
		}
		b.WriteByte(allCharSet[n.Int64()])
	}
	return b.String(), nil
}
func HasUpper(s string) bool {
	for _, r := range s {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
func HasLower(s string) bool {
	for _, r := range s {
		if unicode.IsLower(r) && unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
func HasDigits(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(matchAllCap.ReplaceAllString(snake, "${1}_${2}"))
}
