package main

import (
	"database/sql"
	"net/http"

	"com.ikhsanhaikal.technopartner/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:55000)/technopartner?parseTime=true")

	app := handler.Application{
		DB: db,
	}

	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {})

	r.Get("/home", func(w http.ResponseWriter, r *http.Request) {
		/*
			return { saldo, total pengeluaran, total pemasukan }
		*/
	})

	r.Get("/categories", app.CategoriesList)
	r.Post("/categories", app.CategoriesCreate)
	r.Delete("/categories/{id}", app.CategoriesDelete)

	r.Get("/users/{id}/transactions", app.TransactionsList)
	r.Post("/users/{id}/accounts/{accId}/transactions", app.TransactionsCreate)
	r.Delete("/users/{id}/transactions/{transactionId}", app.TransactionsDelete)

	http.ListenAndServe(":5555", r)
}
