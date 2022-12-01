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

	ID        int64         `bun:"id,pk,autoincrement"`
	Name      string        `bun:",notnull"`
	Kind      AdvertiseKind `bun:",notnull"`
	Provider  string        `bun:",notnull"`
	Country   string        `bun:",notnull"`
	City      string        `bun:",notnull"`
	Street    string
	CreatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
	UpdatedAt time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}
