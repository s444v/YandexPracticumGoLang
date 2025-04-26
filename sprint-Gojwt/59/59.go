package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// Claims хранит идентификатор пользователя в токене
type Claims struct {
	jwt.RegisteredClaims     // базовый тип
	ID                   int // идентификатор пользователя
}

// ValidateToken проверяет валидность JWT-токена и возвращает идентификатор пользователя
func ValidateToken(tokenString string, secretKey string) (bool, int, error) {
	// напишите код функции, который проверяет действителен ли токен и возвращает
	// признак валидности, идентификатор пользователя и ошибку.
	secret := []byte(secretKey)
	claims := Claims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return secret, nil
		})
	if err != nil || !token.Valid {
		return false, 0, err
	}
	return true, claims.ID, err
}

func main() {
	// секретный ключ
	secretKey := "Практикум"

	tokens := []string{
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6MX0.uA0tyxpbQhKk-NubbeT3rn5430WofDOt9FFvh-BYK_M",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Mn0.nDJoURA9MRSaQ7wf6aVsyLqb2RAJdyaersNE7e_RAYg",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJA.eyJJRCI6Nn0.ZfO8KwGlMnzMxeRBRCQ50QVQmUSELZOpd-y5bplYFlW",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6M30.cvXoS7wmwEvuq2DkGFt27olGvrmgSC3qj0I8-_gYQtI",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJA.eyJJRCI6Nn0.ZfO8KwGlMnzMxeRBRCQ50QVQmUSELZOpd-y5bplYFlw",
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6NH0.CkV2Xy3n3It1teNIis-SxiV9SwWhxk-JduGoOK1h4yo",
	}

	for _, token := range tokens {
		isValid, id, err := ValidateToken(token, secretKey)
		if err == nil && isValid {
			fmt.Println(id)
		}
	}
}
