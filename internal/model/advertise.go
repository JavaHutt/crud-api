package model

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/uptrace/bun"
)

// AdvertiseKind is a kind of outdoor advertising
type AdvertiseKind string

const (
	AdvertiseKindBillboard     AdvertiseKind = "billboard"
	AdvertiseKindCitylight     AdvertiseKind = "citylight"
	AdvertiseKindStander       AdvertiseKind = "stander"
	AdvertiseKindLightbox      AdvertiseKind = "lightbox"
	AdvertiseKindBannerStretch AdvertiseKind = "banner_stretch"
	AdvertiseKindPillar        AdvertiseKind = "pillar"
	AdvertiseKindTransition    AdvertiseKind = "transition"
	AdvertiseKindSignboard     AdvertiseKind = "signboard"
	AdvertiseKindAeroman       AdvertiseKind = "aeroman"
	AdvertiseKindNeon          AdvertiseKind = "neon"
)

// Advertise is an outdoor advertise entity
type Advertise struct {
	bun.BaseModel `bun:"table:advertise,alias:a"`

	ID        int64         `json:"id" bun:"id,pk,autoincrement"`
	Name      string        `json:"name" bun:",notnull" validate:"required,alpha"`
	Kind      AdvertiseKind `json:"kind" bun:",notnull" validate:"required,kindenum"`
	Provider  string        `json:"provider" bun:",notnull" validate:"required,alpha"`
	Country   string        `json:"country" bun:",notnull" validate:"required,alpha"`
	City      string        `json:"city" bun:",notnull" validate:"required,alpha"`
	Street    string        `json:"street" validate:"required_unless=Kind transition,alpha"`
	CreatedAt time.Time     `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time     `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}

// I put variable here for caching purposes
var validate = validator.New()

func RegisterValidators() {
	validate.RegisterValidation("kindenum", validateKind)
}

func validateKind(fl validator.FieldLevel) bool {
	kind := fl.Field().String()
	switch kind {
	case string(AdvertiseKindBillboard),
		string(AdvertiseKindCitylight),
		string(AdvertiseKindStander),
		string(AdvertiseKindLightbox),
		string(AdvertiseKindPillar),
		string(AdvertiseKindTransition),
		string(AdvertiseKindSignboard),
		string(AdvertiseKindAeroman),
		string(AdvertiseKindNeon):
		return true
	default:
		return false
	}
}
