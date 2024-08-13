package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"com.ikhsanhaikal.technopartner/sqlcdb"
	"github.com/go-chi/chi/v5"
)

type Category struct {
	Nama string
	Tipe string
}

func (app *Application) CategoriesList(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)
	categories, err := queries.ListCategories(context.Background())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	categoriesJson, _ := json.Marshal(categories)
	w.Write(categoriesJson)
}

func (app *Application) CategoriesCreate(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	c := Category{}

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	result, err := queries.CreateCategory(context.Background(), sqlcdb.CreateCategoryParams{
		Nama: c.Nama,
		Tipe: sqlcdb.CategoriesTipe(c.Tipe),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	id, _ := result.LastInsertId()

	lastInsertedCategory, err := queries.GetCategory(r.Context(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	category, _ := queries.GetCategory(r.Context(), int32(lastInsertedCategory.ID))

	fmt.Printf("category: %+v\n", category)

	w.Write([]byte("ok"))
}

func (app *Application) CategoriesDelete(w http.ResponseWriter, r *http.Request) {
	queries := sqlcdb.New(app.DB)

	stringId := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(stringId, 10, 32)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	category, err := queries.GetCategory(r.Context(), int32(id))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := queries.DeleteCategory(r.Context(), int32(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Printf("category: %+v\n", category)

	w.Write([]byte("ok"))
}
