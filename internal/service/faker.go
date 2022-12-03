package service

import (
	"sync"
	"time"

	"github.com/JavaHutt/crud-api/internal/model"

	"github.com/brianvoe/gofakeit/v6"
)

var statementKinds = []model.QueryStatement{
	model.QueryStatementSelect,
	model.QueryStatementInsert,
	model.QueryStatementUpdate,
	model.QueryStatementDelete,
}

type fakerService struct{}

// NewFakerService is a constructor for faker service
func NewFakerService() fakerService {
	return fakerService{}
}

// Fake generates random query objects, where num is the number of objects
func (svc fakerService) Fake(num int) []model.SlowestQuery {
	faker := gofakeit.New(0)
	now := time.Now().UTC()
	queries := make([]model.SlowestQuery, 0, num)
	var mtx sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < num; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mtx.Lock()
			defer mtx.Unlock()
			queries = append(queries, generateQuery(faker, now))
		}()
	}
	wg.Wait()
	return queries
}

func generateQuery(faker *gofakeit.Faker, now time.Time) model.SlowestQuery {
	statement := generateStatementKind(faker)
	return model.SlowestQuery{
		Query:     faker.Sentence(faker.IntRange(4, 50)),
		Statement: statement,
		TimeSpent: faker.IntRange(1, 1000),
		CreatedAt: faker.DateRange(now.AddDate(-3, 0, 0), now.AddDate(-1, 0, 0)),
		UpdatedAt: faker.DateRange(now.AddDate(-1, 0, 0), now),
	}
}

func generateStatementKind(faker *gofakeit.Faker) model.QueryStatement {
	return statementKinds[faker.Number(0, len(statementKinds)-1)]
}
