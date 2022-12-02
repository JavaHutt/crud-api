package model

import (
	"testing"
)

const (
	name     = "banner"
	provider = "outbrain"
	country  = "Switzerland"
	city     = "Bern"
	street   = "Kramgasse"
)

func TestValidateAdvertise(t *testing.T) {
	RegisterValidators()
	type fields struct {
		ID       int64
		Name     string
		Kind     AdvertiseKind
		Provider string
		Country  string
		City     string
		Street   string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "empty",
			wantErr: true,
		},
		{
			name: "no name",
			fields: fields{
				Kind:     AdvertiseKindBannerStretch,
				Provider: provider,
				Country:  country,
				City:     city,
				Street:   street,
			},
			wantErr: true,
		},
		{
			name: "no kind",
			fields: fields{
				Name:     name,
				Provider: provider,
				Country:  country,
				City:     city,
				Street:   street,
			},
			wantErr: true,
		},
		{
			name: "unknown kind",
			fields: fields{
				Name:     name,
				Kind:     "prebid",
				Provider: provider,
				Country:  country,
				City:     city,
				Street:   street,
			},
			wantErr: true,
		},
		{
			name: "no provider",
			fields: fields{
				Name:    name,
				Kind:    AdvertiseKindAeroman,
				Country: country,
				City:    city,
				Street:  street,
			},
			wantErr: true,
		},
		{
			name: "no country",
			fields: fields{
				Name:     name,
				Kind:     AdvertiseKindAeroman,
				Provider: provider,
				City:     city,
				Street:   street,
			},
			wantErr: true,
		},
		{
			name: "no city",
			fields: fields{
				Name:     name,
				Kind:     AdvertiseKindAeroman,
				Provider: provider,
				Country:  country,
				Street:   street,
			},
			wantErr: true,
		},
		{
			name: "no street",
			fields: fields{
				Name:     name,
				Kind:     AdvertiseKindAeroman,
				Provider: provider,
				Country:  country,
				City:     city,
			},
			wantErr: true,
		},
		{
			name: "no street with transition kind",
			fields: fields{
				Name:     name,
				Kind:     AdvertiseKindTransition,
				Provider: provider,
				Country:  country,
				City:     city,
			},
		},
		{
			name: "success",
			fields: fields{
				Name:     name,
				Kind:     AdvertiseKindAeroman,
				Provider: provider,
				Country:  country,
				City:     city,
				Street:   street,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			a := Advertise{
				ID:       tc.fields.ID,
				Name:     tc.fields.Name,
				Kind:     tc.fields.Kind,
				Provider: tc.fields.Provider,
				Country:  tc.fields.Country,
				City:     tc.fields.City,
				Street:   tc.fields.Street,
			}
			if err := Validate.Struct(a); (err != nil) != tc.wantErr {
				t.Errorf("Advertise.Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
