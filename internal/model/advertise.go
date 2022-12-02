package model

import (
	"time"

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
	Street    string        `json:"street" validate:"required_unless=Kind transition"`
	CreatedAt time.Time     `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time     `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}
