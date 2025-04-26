package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "qw45jk32"

// Claims хранит идентификатор пользователя в токене
type Claims struct {
	ID                   int    // идентификатор пользователя
	Name                 string // идентификатор пользователя
	jwt.RegisteredClaims        // базовый тип
}

// CreateToken создаёт токен
func CreateToken(id int, name string) (string, error) {
	// напишите код функции, которая создаёт и возвращает JWT токен
	claim := Claims{ID: id, Name: name}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil

}

// ValidateToken проверяет валидность JWT-токена и возвращает идентификатор пользователя
func ValidateToken(tokenString string, claims *Claims) bool {
	// парсим токен
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}

func main() {
	for i, name := range []string{"Андрей", "Олег", "Наталья"} {
		token, err := CreateToken(i+1, name)
		if err != nil {
			return
		}
		var claim Claims
		if ValidateToken(token, &claim) {
			fmt.Println(claim.ID, claim.Name)
		}
	}
}
