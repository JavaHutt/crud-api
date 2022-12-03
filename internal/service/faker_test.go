package service

import (
	"testing"
	"time"
	"unicode/utf8"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestFake(t *testing.T) {
	svc := NewFakerService()
	queries := svc.Fake(10000)
	require.Len(t, queries, 10000)
}

func TestGenerateQuery(t *testing.T) {
	faker := gofakeit.New(0)
	now := time.Now().UTC()
	for i := 0; i < 1000; i++ {
		ad := generateQuery(faker, now)
		require.LessOrEqual(t, ad.CreatedAt, ad.UpdatedAt)
	}
}

func FuzzGenerateStatementKind(f *testing.F) {
	testcases := []int64{0, 10, time.Now().UnixNano()}
	for _, tc := range testcases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, seed int64) {
		faker := gofakeit.New(seed)
		kind := generateStatementKind(faker)

		require.Contains(t, statementKinds, kind)
		if !utf8.ValidString(string(kind)) {
			t.Errorf("Kind should be valid string %q", kind)
		}
	})
}
