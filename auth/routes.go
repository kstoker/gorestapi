package auth

import (
	"encoding/json"
	"net/http"
//	"time"
	"io/ioutil"
	"io"
	"golang.org/x/crypto/bcrypt"
//	"strings"

	"github.com/gorilla/mux"
//	"github.com/gorilla/context"
	"github.com/dgrijalva/jwt-go"
//	"github.com/mitchellh/mapstructure"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gorestapi/database"
	logger "gorestapi/logger"
	config "gorestapi/config"
)

// Create a struct that models the structure of a user in the request body
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Exception struct {
	Message string `json:"message"`
}

// Set user table name to be `auth_user`
func (User) TableName() string {
	return "auth_user"
}

func AddToRouter(r *mux.Router) {
	// Route handlers and routes
	r.HandleFunc("/api/createtoken", CreateToken).Methods("POST")
}

func CreateToken(w http.ResponseWriter, r *http.Request){
	// Parse and decode the request body into a new `Credentials` instance
	var creds Credentials
	// CreateBodyLimit is ingesteld op 1 Mb als maximale omvang van de body
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, config.CreateBodyLimit))
    if err != nil {
		logger.Log.Println(err)
        w.WriteHeader(http.StatusBadRequest)
		return
    }
    if err := r.Body.Close(); err != nil {
		logger.Log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
		return
    }
    if err := json.Unmarshal(body, &creds); err != nil {
		logger.Log.Println(err)
        w.WriteHeader(422) // unprocessable entity
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        if err := json.NewEncoder(w).Encode(err); err != nil {
			logger.Log.Println(err)
        }
		return
	}
	// Get the existing entry present in the database for the given username
	var user User
	db := database.GetDB()
	db.Where("username = ?", creds.Username).First(&user)
	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// If we reach this point user/password was correct
	// nu kunnen we een token genereren
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"password": user.Password,
	})
	tokenString, err := token.SignedString([]byte(config.SecurityTokenSecret))
	if err != nil {
		logger.Log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(tokenString)
}
