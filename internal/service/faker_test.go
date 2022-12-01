package service

import (
	"testing"
	"time"
	"unicode/utf8"

	"github.com/JavaHutt/crud-api/internal/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestFake(t *testing.T) {
	svc := NewFakerService()
	ads := svc.Fake(10000)
	require.Len(t, ads, 10000)
}

func TestGenerateAdvertise(t *testing.T) {
	faker := gofakeit.New(0)
	now := time.Now().UTC()
	for i := 0; i < 1000; i++ {
		ad := generateAdvertise(faker, now)
		require.LessOrEqual(t, ad.CreatedAt, ad.UpdatedAt)
		if ad.Kind == model.AdvertiseKindTransition {
			require.Empty(t, ad.Street)
		}
	}
}

func FuzzGenerateAdvertiseKind(f *testing.F) {
	testcases := []int64{0, 10, time.Now().UnixNano()}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, seed int64) {
		faker := gofakeit.New(seed)
		kind := generateAdvertiseKind(faker)

		require.Contains(t, adKinds, kind)
		if !utf8.ValidString(string(kind)) {
			t.Errorf("Kind should be valid string %q", kind)
		}
	})
}
func FuzzGenerateProvider(f *testing.F) {
	testcases := []int64{0, 10, time.Now().UnixNano()}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, seed int64) {
		faker := gofakeit.New(seed)
		provider := generateProvider(faker)

		require.Contains(t, providers, provider)
		if !utf8.ValidString(provider) {
			t.Errorf("Provider should be valid string %q", provider)
		}
	})
}
