package handler

import (
	"encoding/json"
	"fmt"
	"gintraining/database"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	Router *http.ServeMux
	DB     database.IDatabase
}

func (h *Handler) routes() {
	h.Router.HandleFunc("/sign-in", h.SignIn)
	h.Router.HandleFunc("/sign-up", h.SignUp)
	h.Router.HandleFunc("/me", h.CheckMe)
}
func (h *Handler) CheckMe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		bearer := r.Header.Get("Authorization")
		token, err := Validate(bearer)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		user := token.Claims.(*Token)
		json.NewEncoder(w).Encode(fmt.Sprintf("User: %s", user.Email))
	}
}
func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var user *User
		json.NewDecoder(r.Body).Decode(&user)
		if h.DB.CheckOnLogin((*database.User)(user)) {
			var tokenClaim = &Token{
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Minute).Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaim)
			tokenString, _ := token.SignedString(InitJWT())
			json.NewEncoder(w).Encode(tokenString)
			w.Write([]byte("email and pass are matching the account"))
		} else {
			w.Write([]byte("incorrect mail or password"))
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func Validate(bearer string) (*jwt.Token, error) {
	tokenString := strings.Split(bearer, " ")[1]
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return InitJWT(), nil
	})
	return token, err
}
func (h *Handler) SignUp(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Write([]byte("Please, enter your login and password, tho there is no form atm B)"))
	case "POST":
		var user *User
		json.NewDecoder(r.Body).Decode(&user)
		if !h.DB.CheckForExisting((*database.User)(user)) {
			h.DB.DBAddUser((*database.User)(user))
			w.Write([]byte("successfully registered"))
		} else {
			w.Write([]byte("Account with this email or username already exists"))
		}
	}
}
func InitHandler() *Handler {
	h := &Handler{
		Router: http.NewServeMux(),
		DB:     database.InitDB("sqlite3"),
	}
	h.routes()
	h.DB.DBCreateTables()
	return &Handler{
		Router: h.Router,
		DB:     h.DB,
	}
}
