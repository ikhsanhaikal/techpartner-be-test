package handler

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
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

	resp, _ := json.Marshal(transactions)
	fmt.Printf("transactions: %+v\n", transactions)

	w.Header().Set("content-type", "application/json")
	w.Write(resp)
}

func (app *Application) TransactionsCreate(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringUserId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(stringUserId, 10, 32)
	userId := int32(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	stringRekId := chi.URLParam(r, "accId")
	accId, _ := strconv.ParseInt(stringRekId, 10, 32)
	rekId := int32(accId)

	t := Transaction{}

	// fmt.Printf("hi 1\n")
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// fmt.Printf("hi 2\n")
	deskripsi := sql.NullString{}

	accounts, err := queries.GetAccounts(r.Context(), sqlcdb.GetAccountsParams{
		ID:     rekId,
		UserID: userId,
	})

	fmt.Printf("accounts: %+v\n", accounts)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if t.Deskripsi != "" {
		deskripsi.String = t.Deskripsi
		deskripsi.Valid = true
	}

	tx, err := app.DB.Begin()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Printf("hi 3\n")
	qtx := queries.WithTx(tx)

	category, err := qtx.GetCategory(r.Context(), t.CategoryId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	// fmt.Printf("category: %+v\n", category)
	// fmt.Printf("tipe: %s\n", tipe)

	if category.Tipe == "pemasukan" {
		err = qtx.Deposit(r.Context(), sqlcdb.DepositParams{
			Saldo: t.Nominal,
			ID:    rekId,
		})
	} else {
		err = qtx.Withdraw(r.Context(), sqlcdb.WithdrawParams{
			Saldo: t.Nominal,
			ID:    rekId,
		})
	}

	// fmt.Printf("hi 4\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	// fmt.Printf("hi 5\n")
	transaction, err := qtx.CreateTransaction(context.Background(), sqlcdb.CreateTransactionParams{
		Nominal:    t.Nominal,
		KategoriID: t.CategoryId,
		Deskripsi:  deskripsi,
		UserID:     userId,
		RekID:      rekId,
	})

	// fmt.Printf("hi 6\n")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		tx.Rollback()
		return
	}

	// fmt.Printf("hi 7\n")
	tx.Commit()

	// fmt.Printf("hi 8\n")
	fmt.Printf("\ntransaction: %+v\n", transaction)

	w.Write([]byte("ok"))
}

func (app *Application) TransactionsDelete(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringId := chi.URLParam(r, "transactionId")
	id, err := strconv.ParseInt(stringId, 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	transaction, err := queries.GetTransactions(r.Context(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := queries.DeleteTransaction(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("transaction: %+v\n", transaction)

	w.Write([]byte("ok"))
}
