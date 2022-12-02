package model

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func RegisterValidators() {
	_ = Validate.RegisterValidation("kindenum", validateKind)
}

func validateKind(fl validator.FieldLevel) bool {
	kind := fl.Field().String()
	switch AdvertiseKind(kind) {
	case AdvertiseKindBillboard,
		AdvertiseKindCitylight,
		AdvertiseKindStander,
		AdvertiseKindLightbox,
		AdvertiseKindPillar,
		AdvertiseKindTransition,
		AdvertiseKindSignboard,
		AdvertiseKindAeroman,
		AdvertiseKindNeon:
		return true
	default:
		return false
	}
}
