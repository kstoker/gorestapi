package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"

	config "gorestapi/config"
	utils "gorestapi/utils"
)

// AuthMiddleware : check if there is a token and if it is correct
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		authorizationHeader := req.Header.Get("Authorization")
		if authorizationHeader != "" {
			bearerToken := strings.Split(authorizationHeader, " ") // first part contains "bearer "
			if len(bearerToken) == 2 {
				token, error := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("An error has occured")
					}
					return []byte(config.SecurityTokenSecret), nil
				})
				if error != nil {
					utils.WriteJSONMessage(w, http.StatusBadRequest, error.Error())
					return
				}
				if token.Valid {
					context.Set(req, "decoded", token.Claims)
					next(w, req)
				} else {
					utils.WriteJSONMessage(w, http.StatusUnauthorized, "Invalid token")
				}
			} else {
				utils.WriteJSONMessage(w, http.StatusBadRequest, "Authorization header is invalid")
			}
		} else {
			utils.WriteJSONMessage(w, http.StatusBadRequest, "An authorization header is required")
		}
	})
}
