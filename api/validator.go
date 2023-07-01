package api

import (
	util "github.com/Hedi-zouaoui/go_project/utils"
	"github.com/go-playground/validator"
)

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}