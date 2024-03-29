package main

import (
	"context"
	"encoding/json"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"mihaly.codes/cart-recommendation-engine/database"
	"mihaly.codes/cart-recommendation-engine/recommender"
)

type SearchProductsResult struct {
	Items []database.SearchProductsRow
	Count int
}

func ChunkList[T any](lst []T, chunkSize int) [][]T {
	var chunks [][]T

	for i := 0; i < len(lst); i += chunkSize {
		end := i + chunkSize
		if end > len(lst) {
			end = len(lst)
		}
		chunks = append(chunks, lst[i:end])
	}

	return chunks
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)

	ctx := context.Background()

	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		slog.Error("failed to connect to PostgreSQL", err)
		os.Exit(1)
	}
	defer conn.Close()

	queries := database.New(conn)

	enableNeo4j := os.Getenv("NEO4J_USER") != "" ||
		os.Getenv("NEO4J_PASSWORD") != "" ||
		os.Getenv("NEO4J_URI") != ""

	var rec recommender.Recommender
	if enableNeo4j {
		rec, err = recommender.NewNeo4jRecommender(
			os.Getenv("NEO4J_USER"),
			os.Getenv("NEO4J_PASSWORD"),
			os.Getenv("NEO4J_URI"),
			func(u uuid.UUID) error { return queries.DeleteCart(ctx, u) },
			ctx,
		)
	} else {
		rec, err = recommender.NewSqlRecommender(
			*queries,
			ctx,
		)

	}

	if err != nil {
		slog.Error("failed to create recommender", err)
		os.Exit(1)
	}
	defer rec.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
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

		tmpl := template.Must(template.ParseFiles("./templates/index.html"))

		data := map[string][]database.ListProductsRow{
			"Product": products,
		}
		tmpl.Execute(w, data)
	})

	http.HandleFunc("GET /products", func(w http.ResponseWriter, req *http.Request) {
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

	http.HandleFunc("GET /products/search", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		searchValue := req.URL.Query().Get("q")

		skipRaw := req.URL.Query().Get("skip")
		var skip int32
		if skipRaw != "" {
			skipParsed, err := strconv.Atoi(skipRaw)
			if err != nil {
				http.Error(w, "failed to parse skip as int", http.StatusBadRequest)
				slog.Error("failed to parse skip as int", "skipRaw", skipRaw, "err", err)
				return
			}
			skip = int32(skipParsed)
		} else {
			skip = 0
		}

		takeRaw := req.URL.Query().Get("take")
		var take int32
		if takeRaw != "" {
			takeParsed, err := strconv.Atoi(takeRaw)
			if err != nil {
				http.Error(w, "failed to parse take as int", http.StatusBadRequest)
				slog.Error("failed to parse take as int", "takeRaw", takeRaw, "err", err)
				return
			}
			take = int32(takeParsed)
		} else {
			take = 25
		}

		products, err := queries.SearchProducts(ctx, database.SearchProductsParams{SearchValue: searchValue, Skip: skip, Take: take})
		if err != nil {
			http.Error(w, "failed to fetch products", http.StatusInternalServerError)
			slog.Error("failed to fetch products", err)
			return
		}
		responseJSON, err := json.Marshal(SearchProductsResult{
			Items: products,
			Count: len(products),
		})
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("GET /products/recommended", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		var cartId uuid.UUID
		if err := cartId.Scan(req.URL.Query().Get("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		recommendedProductIds, err := rec.GetRecommendedItems(cartId)

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

	http.HandleFunc("PUT /carts/{cartId}/items/{itemId}", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		itemId := req.PathValue("itemId")
		var cartId uuid.UUID
		if err := cartId.Scan(req.PathValue("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		tx, err := conn.Begin(ctx)

		if err != nil {
			http.Error(w, "failed begin tx", http.StatusInternalServerError)
			slog.Error("failed begin tx", err)
			return
		}
		defer tx.Rollback(ctx)
		qtx := queries.WithTx(tx)

		err = qtx.PutCart(ctx, cartId)

		if err != nil {
			http.Error(w, "failed to put cart", http.StatusInternalServerError)
			slog.Error("failed to put cart", err)
			return
		}

		err = qtx.PutCartItem(ctx, database.PutCartItemParams{
			CartID:    cartId,
			ProductID: itemId,
		})

		if err != nil {
			http.Error(w, "failed to put cart item", http.StatusInternalServerError)
			slog.Error("failed to put cart item", err)
			return
		}

		tx.Commit(ctx)

		w.WriteHeader(http.StatusNoContent)
	})

	http.HandleFunc("DELETE /carts/{cartId}/items/{itemId}", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		itemId := req.PathValue("itemId")
		var cartId uuid.UUID
		if err := cartId.Scan(req.PathValue("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		err := queries.DeleteCartItem(ctx, database.DeleteCartItemParams{
			CartID:    cartId,
			ProductID: itemId,
		})

		if err != nil {
			http.Error(w, "failed to delete cart", http.StatusInternalServerError)
			slog.Error("failed to delete cart", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})

	http.HandleFunc("GET /carts/{cartId}/items", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		var cartId uuid.UUID
		if err := cartId.Scan(req.PathValue("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		cart, err := queries.GetCartItems(ctx, cartId)

		if err != nil {
			http.Error(w, "failed to get cart items", http.StatusInternalServerError)
			slog.Error("failed to get cart items", err)
			return
		}

		responseJSON, err := json.Marshal(cart)
		if err != nil {
			http.Error(w, "failed to marshal response", http.StatusInternalServerError)
			slog.Error("failed to marshal response", err)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseJSON)
	})

	http.HandleFunc("POST /carts/{cartId}/checkout", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		var cartId uuid.UUID
		if err := cartId.Scan(req.PathValue("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		cartItems, err := queries.GetCartItems(ctx, cartId)

		if err != nil {
			http.Error(w, "failed to get cart items", http.StatusInternalServerError)
			slog.Error("failed to get cart items", err)
			return
		}

		if len(cartItems) == 0 {
			http.Error(w, "cart has no items", http.StatusBadRequest)
			slog.Warn("cart has no items", err)
			return
		}

		var cartItemIds []string

		for _, cartItem := range cartItems {
			cartItemIds = append(cartItemIds, cartItem.ID)
		}

		err = rec.AddOrder(cartId, cartItemIds)

		w.WriteHeader(http.StatusNoContent)
	})

	port := ":8090"
	log.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
