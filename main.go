package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"

	"mihaly.codes/shopping-cart-recommendation-engine/database"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL: %v", err)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	ctx1 := context.Background()
	auth := neo4j.BasicAuth(os.Getenv("NEO4J_USER"), os.Getenv("NEO4J_PASSWORD"), "")
	driver, err := neo4j.NewDriverWithContext(os.Getenv("NEO4J_URI"), auth)
	defer driver.Close(ctx1)

	err = driver.VerifyConnectivity(ctx1)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/products", func(w http.ResponseWriter, req *http.Request) {
		products, err := queries.ListProducts(ctx, database.ListProductsParams{Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/search", func(w http.ResponseWriter, req *http.Request) {
		searchValue := req.URL.Query().Get("q")

		products, err := queries.SearchProducts(ctx, database.SearchProductsParams{SearchValue: searchValue, Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/recommended", func(w http.ResponseWriter, req *http.Request) {
		result, _ := neo4j.ExecuteQuery(ctx, driver,
			"MATCH (p:Product) RETURN p.productId AS productId LIMIT 50",
			map[string]any{}, neo4j.EagerResultTransformer,
			neo4j.ExecuteQueryWithDatabase("neo4j"))

		recommendedProductIds := []string{}
		for _, record := range result.Records {
			productId, _ := record.Get("productId")
			recommendedProductIds = append(recommendedProductIds, productId.(string))
		}

		products, err := queries.GetProductsByIds(ctx, recommendedProductIds)
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	port := ":8090"
	log.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
