package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"com.ikhsanhaikal.technopartner/sqlcdb"
	"github.com/go-chi/chi/v5"
)

type Transaction struct {
	CategoryId int32  `json:"kategori"`
	Nominal    string `json:"nominal"`
	Deskripsi  string `json:"deskripsi"`
}

func (app *Application) TransactionsList(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(stringId, 10, 32)

	startStr := chi.URLParam(r, "start")
	endStr := chi.URLParam(r, "end")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var transactions []sqlcdb.Transaction

	if startStr == "" || endStr == "" {
		transactions, err = queries.ListTransactionsByUserIdDefault(context.Background(), int32(id))
	} else {
		layoutFormat := "2006-01-02 15:04:05"
		start, _ := time.Parse(layoutFormat, startStr)
		end, _ := time.Parse(layoutFormat, endStr)

		transactions, err = queries.ListTransactionsByUserIdRange(context.Background(), sqlcdb.ListTransactionsByUserIdRangeParams{
			FromCreatedAt: start,
			ToCreatedAt:   end,
			UserID:        int32(id),
		})
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("transactions: %+v\n", transactions)

	w.Write([]byte("ok"))
}

func (app *Application) TransactionsCreate(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(stringId, 10, 32)
	userId := int32(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	t := Transaction{}

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	deskripsi := sql.NullString{}

	if t.Deskripsi != "" {
		deskripsi.String = t.Deskripsi
		deskripsi.Valid = true
	}

	categories, err := queries.CreateTransaction(context.Background(), sqlcdb.CreateTransactionParams{
		Nominal:    t.Nominal,
		KategoriID: t.CategoryId,
		Deskripsi:  deskripsi,
		UserID:     userId,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("\ncategories: %+v\n", categories)

	w.Write([]byte("ok"))
}

func (app *Application) TransactionsDelete(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringId := chi.URLParam(r, "transactionId")
	id, err := strconv.ParseInt(stringId, 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	transaction, err := queries.GetTransactions(r.Context(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := queries.DeleteTransaction(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("transaction: %+v\n", transaction)

	w.Write([]byte("ok"))
}
