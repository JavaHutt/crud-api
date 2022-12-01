package main

import (
	"context"
	"fmt"
	"log"

	"github.com/JavaHutt/crud-api/internal/repository"

	"github.com/uptrace/bun"
)

func main() {
	ctx := context.Background()
	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var num []int
	err = db.NewRaw(
		"SELECT id FROM ? LIMIT ?",
		bun.Ident("test_table"), 100,
	).Scan(ctx, &num)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
}
