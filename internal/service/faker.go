package service

import (
	"sync"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

var adKinds = []model.AdvertiseKind{
	model.AdvertiseKindBillboard,
	model.AdvertiseKindCitylight,
	model.AdvertiseKindStander,
	model.AdvertiseKindLightbox,
	model.AdvertiseKindBannerStretch,
	model.AdvertiseKindPillar,
	model.AdvertiseKindTransition,
	model.AdvertiseKindSignboard,
	model.AdvertiseKindAeroman,
	model.AdvertiseKindNeon,
}

var providers = []string{"adblock", "outbrain", "plista", "affiliate", "appnext", "bizzclick", "mcn",
	"mobupps", "nativex", "plista", "smartyads", "strossle"}

type fakerService struct{}

// NewFakerService is a constructor for faker service
func NewFakerService() fakerService {
	return fakerService{}
}

// Fake generates random advertise objects, where num is the number of objects
func (svc fakerService) Fake(num int) []model.Advertise {
	faker := gofakeit.New(0)
	now := time.Now().UTC()
	ads := make([]model.Advertise, 0, num)
	var mtx sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mtx.Lock()
			defer mtx.Unlock()
			ads = append(ads, generateAdvertise(faker, now))
		}()
	}
	wg.Wait()
	return ads
}

func generateAdvertise(faker *gofakeit.Faker, now time.Time) model.Advertise {
	kind := generateAdvertiseKind(faker)
	var street string
	if kind != model.AdvertiseKindTransition {
		street = faker.Street()
	}
	return model.Advertise{
		Name:      faker.Noun(),
		Kind:      kind,
		Provider:  generateProvider(faker),
		Country:   faker.Country(),
		City:      faker.City(),
		Street:    street,
		CreatedAt: faker.DateRange(now.AddDate(-3, 0, 0), now.AddDate(-1, 0, 0)),
		UpdatedAt: faker.DateRange(now.AddDate(-1, 0, 0), now),
	}
}

func generateAdvertiseKind(faker *gofakeit.Faker) model.AdvertiseKind {
	return adKinds[faker.Number(0, len(adKinds)-1)]
}

func generateProvider(faker *gofakeit.Faker) string {
	return providers[faker.Number(0, len(providers)-1)]
}
