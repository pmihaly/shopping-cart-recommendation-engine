package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"

	"mihaly.codes/shopping-cart-recommendation-engine/v2/database"
)

func run() error {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, "user=postgres password=hunter2 dbname=shopping-cart-recommendation-engine")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	products, err := queries.ListProducts(ctx)
	if err != nil {
		return err
	}
	log.Println(products)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
