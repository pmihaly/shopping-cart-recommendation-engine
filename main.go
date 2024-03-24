package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"

	"mihaly.codes/cart-recommendation-engine/database"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		slog.Error("failed to connect to PostgreSQL", err)
		os.Exit(1)
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	recommender, err := NewExtraItemRecommender(
		os.Getenv("NEO4J_USER"),
		os.Getenv("NEO4J_PASSWORD"),
		os.Getenv("NEO4J_URI"),
		ctx,
	)
	if err != nil {
		slog.Error("failed to create recommender", err)
		os.Exit(1)
	}
	defer recommender.Close()

	http.HandleFunc("/products", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		products, err := queries.ListProducts(ctx, database.ListProductsParams{Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			slog.Error("failed to fetch products", err)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/search", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		searchValue := req.URL.Query().Get("q")

		products, err := queries.SearchProducts(ctx, database.SearchProductsParams{SearchValue: searchValue, Skip: 0, Take: 25})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			slog.Error("failed to fetch products", err)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("/products/recommended", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		cartId := req.URL.Query().Get("cartId")

		recommendedProductIds, err := recommender.GetRecommendedExtraItems(cartId)
		if err != nil {
			http.Error(w, "failed to recommend products", http.StatusInternalServerError)
			return
		}

		products, err := queries.GetProductsByIds(ctx, recommendedProductIds)
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			slog.Error("failed to fetch products", err)
			return
		}
		responseJSON, err := json.Marshal(products)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response", err)
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
