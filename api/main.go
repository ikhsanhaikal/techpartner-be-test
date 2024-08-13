package main

import (
	"database/sql"
	"net/http"

	"com.ikhsanhaikal.technopartner/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:55000)/technopartner?parseTime=true")

	app := handler.Application{
		DB:        db,
		SecretKey: []byte("some-highclassified-key"),
	}

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", app.Login)

	r.Group(func(r chi.Router) {
		r.Use(func(h http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				tokenString := r.Header.Get("Authorization")

				if tokenString == "" {
					http.Error(w, "unauthorized", http.StatusUnauthorized)
					return
				}

				tokenString = tokenString[len("Bearer "):]

				token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
					return app.SecretKey, nil
				})

				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				if !token.Valid {
					http.Error(w, "invalid token", http.StatusInternalServerError)
					return
				}

				h.ServeHTTP(w, r)
			})
		})
		r.Get("/categories", app.CategoriesList)
		r.Post("/categories", app.CategoriesCreate)
		r.Delete("/categories/{id}", app.CategoriesDelete)

		r.Get("/users/{id}/transactions", app.TransactionsList)
		r.Post("/users/{id}/accounts/{accId}/transactions", app.TransactionsCreate)
		// r.Delete("/users/{id}/transactions/{transactionId}", app.TransactionsDelete)

	})

	http.ListenAndServe(":5555", r)
}
