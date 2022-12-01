package model

import "time"

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
	Name      string
	Kind      AdvertiseKind
	Provider  string
	Country   string
	City      string
	Street    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
