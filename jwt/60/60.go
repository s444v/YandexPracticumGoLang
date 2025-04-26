package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var (
	secret = []byte("gBElG5NThZSye")
)

type Claims struct {
	Name string
	jwt.RegisteredClaims
}

func verifyUser(token string) (string, bool) {
	var claims Claims
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		// проверяем алгоритм подписи
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return secret, nil
	})
	if err != nil || !jwtToken.Valid {
		return ``, false
	}
	return claims.Name, true
}

func main() {
	address := ":8080"

	mux := http.NewServeMux()
	mux.HandleFunc("/secret/", func(w http.ResponseWriter, req *http.Request) {
		// получаем http заголовок вида 'Bearer {jwt}'
		authHeaderValue := req.Header.Get("Authorization")
		// проверяем доступы
		if authHeaderValue != "" {
			bearerToken := strings.Split(authHeaderValue, " ")
			if len(bearerToken) == 2 {
				if username, ok := verifyUser(bearerToken[1]); ok {
					fmt.Fprintf(w, "Добро пожаловать, %s!\n", username)
					return
				}
			}
		}
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	})

	srv := &http.Server{
		Addr:    address,
		Handler: mux,
	}

	fmt.Printf("Starting server on %s\n", address)
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Printf("failed to listen and serve: %s\n", err)
	}
}
