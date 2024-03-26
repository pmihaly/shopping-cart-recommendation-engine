package main

import (
	"context"
	"encoding/json"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

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

	http.HandleFunc("/carts/{cartId}/items/{itemId}", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		itemId := req.PathValue("itemId")
		var cartId pgtype.UUID
		if err := cartId.Scan(req.PathValue("cartId")); err != nil {
			http.Error(w, "failed to parse UUID", http.StatusInternalServerError)
			slog.Error("failed to parse UUID", err, "cartId", req.PathValue("cartId"))
			return
		}

		if req.Method == "PUT" {
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
		}

		if req.Method == "DELETE" {
			err := queries.DeleteCartItem(ctx, database.DeleteCartItemParams{
				CartID:    cartId,
				ProductID: itemId,
			})

			if err != nil {
				http.Error(w, "failed to delete cart", http.StatusInternalServerError)
				slog.Error("failed to delete cart", err)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	})

	http.HandleFunc("/carts/{cartId}/items", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		var cartId pgtype.UUID
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

	http.HandleFunc("/carts/{cartId}/checkout", func(w http.ResponseWriter, req *http.Request) {
		logger.Info(
			"incoming request",
			"method", req.Method,
			"path", req.URL.RequestURI(),
			"user_agent", req.UserAgent(),
		)

		if req.Method != "POST" {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var cartId pgtype.UUID
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
			slog.Warn("cart has no items", "cartId", cartId, "cartItems", cartItems)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}

		var cartItemIds []string

		for _, cartItem := range cartItems {
			cartItemIds = append(cartItemIds, cartItem.ID)
		}

		err = queries.DeleteCart(ctx, cartId)
		if err != nil {
			http.Error(w, "failed to delete cart", http.StatusInternalServerError)
			slog.Error("failed to delete cart", err)
			return
		}

		err = recommender.AddOrder(req.PathValue("cartId"), cartItemIds)

		w.WriteHeader(http.StatusNoContent)
	})

	port := ":8090"
	log.Printf("Server started at %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
