package model

import (
	"testing"
)

func TestAdvertiseValidate(t *testing.T) {
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
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
				Street:   "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "no kind",
			fields: fields{
				Name:     "Banner",
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
				Street:   "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "unknown kind",
			fields: fields{
				Name:     "Banner",
				Kind:     "prebid",
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
				Street:   "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "no provider",
			fields: fields{
				Name:    "Banner",
				Kind:    AdvertiseKindAeroman,
				Country: "Switzerland",
				City:    "Bern",
				Street:  "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "no country",
			fields: fields{
				Name:     "Banner",
				Kind:     AdvertiseKindAeroman,
				Provider: "outbrain",
				City:     "Bern",
				Street:   "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "no city",
			fields: fields{
				Name:     "Banner",
				Kind:     AdvertiseKindAeroman,
				Provider: "outbrain",
				Country:  "Switzerland",
				Street:   "Kramgasse",
			},
			wantErr: true,
		},
		{
			name: "no street",
			fields: fields{
				Name:     "Banner",
				Kind:     AdvertiseKindAeroman,
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
			},
			wantErr: true,
		},
		{
			name: "no street with transition kind",
			fields: fields{
				Name:     "Banner",
				Kind:     AdvertiseKindTransition,
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
			},
			wantErr: true,
		},
		{
			name: "success",
			fields: fields{
				Name:     "Banner",
				Kind:     AdvertiseKindAeroman,
				Provider: "outbrain",
				Country:  "Switzerland",
				City:     "Bern",
				Street:   "Kramgasse",
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
			if err := validate.Struct(a); (err != nil) != tc.wantErr {
				t.Errorf("Advertise.Validate() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
