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
	Name      string        `json:"name" bun:",notnull"`
	Kind      AdvertiseKind `json:"kind" bun:",notnull"`
	Provider  string        `json:"provider" bun:",notnull"`
	Country   string        `json:"country" bun:",notnull"`
	City      string        `json:"city" bun:",notnull"`
	Street    string        `json:"street"`
	CreatedAt time.Time     `json:"created_at" bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time     `json:"updated_at" bun:",nullzero,notnull,default:current_timestamp"`
}
