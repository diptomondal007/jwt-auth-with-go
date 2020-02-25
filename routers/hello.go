package routers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/chi"
	"jwtauthwithgo/handlers"
	"net/http"
)

type Exception struct {
	Message string `json:"message"`
}

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type","application/json")

		if r.Header["Authorization"] != nil {

			token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret"), nil
			})

			if err != nil {
				json.NewEncoder(w).Encode(Exception{Message: err.Error()})
				return
			}

			if token.Valid {
				next.ServeHTTP(w, r)
			} else {

				json.NewEncoder(w).Encode(Exception{Message: "Invalid authorization token"})
			}
		} else {
			json.NewEncoder(w).Encode(Exception{Message: "An authorization header is required"})
		}

	})
}

func SetHelloRoutes(mux *chi.Mux) *chi.Mux {
	mux.With(ValidateMiddleware).Get("/test/hello", handlers.HelloHandler)
	return mux
}
