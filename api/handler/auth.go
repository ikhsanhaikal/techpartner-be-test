package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"com.ikhsanhaikal.technopartner/sqlcdb"
	"github.com/golang-jwt/jwt/v5"
)

type userDto struct {
	Email    string
	Password string
}

func (app *Application) Login(w http.ResponseWriter, r *http.Request) {

	queries := sqlcdb.New(app.DB)

	var userDto userDto
	json.NewDecoder(r.Body).Decode(&userDto)

	user, err := queries.GetUserByEmail(r.Context(), userDto.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Password != userDto.Password { // TODO: hashing password
		http.Error(w, "invalid username / password", http.StatusBadRequest)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": user.Email,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, _ := token.SignedString(app.SecretKey)
	fmt.Fprint(w, tokenString)
}
