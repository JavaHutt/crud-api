package model

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func RegisterValidators() {
	_ = Validate.RegisterValidation("statementenum", validateStatement)
}

func validateStatement(fl validator.FieldLevel) bool {
	statement := fl.Field().String()
	switch QueryStatement(statement) {
	case QueryStatementSelect,
		QueryStatementInsert,
		QueryStatementUpdate,
		QueryStatementDelete:
		return true
	default:
		return false
	}
}
