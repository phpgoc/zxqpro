package request

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ComplexPasswordValidator = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	// 定义正则表达式来匹配不同类型的字符
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	// 计算字符类型的数量
	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasDigit {
		count++
	}
	if hasSpecial {
		count++
	}

	return count >= 3
}
