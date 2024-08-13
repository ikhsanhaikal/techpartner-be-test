package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"com.ikhsanhaikal.technopartner/sqlcdb"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:55000)/technopartner")

	if err != nil {
		panic(err)
	}

	queries := sqlcdb.New(db)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	r.Get("/categories", func(w http.ResponseWriter, r *http.Request) {

		categories, err := queries.ListCategories(context.Background())

		if err != nil {
			msgJson, _ := json.Marshal(map[string]string{
				"message": "not ok",
				"err":     fmt.Sprintf("%s", err.Error()),
			})
			w.Write(msgJson)
			return
		}

		fmt.Printf("%+v\n", categories)

		w.Header().Set("content-type", "application/json")

		categoriesJson, _ := json.Marshal(categories)

		fmt.Printf("json: %s\n", categoriesJson)

		// msgJson, _ := json.Marshal(map[string]interface{}{
		// 	"message":    "ok",
		// 	"categories": categoriesJson,
		// })
		w.Write(categoriesJson)
	})

	r.Get("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
		strId := chi.URLParam(r, "userId")

		userId, _ := strconv.ParseInt(strId, 10, 32)

		user, err := queries.GetUser(context.Background(), int32(userId))

		if err != nil {
			w.Write([]byte("ok"))
		}

		data, _ := json.Marshal(user)
		w.Header().Set("content-type", "application/json")
		w.Write(data)
	})

	http.ListenAndServe(":5555", r)
}
