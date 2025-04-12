package main

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func main() {
	// для примера возьмём токен, подписанный при помощи секретного ключа secretKey
	token := "eeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbiI6ImFsaWNlIiwicm9sZXMiOlsicmVhZGVyIiwid3JpdGVyIl19.Fppwzfqcl8oxEqcQmDesbzJ77DUDdsqLhh3emlKmv3s"
	secretKey := []byte("my_secret_key")

	// парсим токен
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Printf("failed to parse token: %s\n", err)
		return
	}
	if !jwtToken.Valid {
		fmt.Printf("token is invalid")
		return
	}
	// приводим поле Claims к типу jwt.MapClaims
	res, ok := jwtToken.Claims.(jwt.MapClaims)
	// обязательно используем второе возвращаемое значение ok и проверяем его, потому что
	// если Сlaims вдруг окажется другого типа, мы получим панику
	if !ok {
		fmt.Printf("failed to typecast to jwt.MapCalims")
		return
	}
	// Так как jwt.Claims — словарь вида map[string]inteface{}, используем синтаксис получения
	// значения по ключу. Получаем значение ключа "login" и "roles"
	loginRaw := res["login"]
	rolesRaw := res["roles"]
	// loginRaw — интерфейс, так как тип значения в jwt.Claims — интерфейс.
	// Чтобы получить строку, нужно снова сделать приведение типа к строке.
	login, ok := loginRaw.(string)
	if !ok {
		fmt.Printf("failed to typecast to string")
		return
	}
	// обратите внимание, что при создании ролей мы указывали тип []string,
	// однако тут приводим к []inteface{}
	// так происходит, потому что json не строго типизированный,
	// из-за чего при парсинге нельзя точно определить тип слайса.
	roles, ok := rolesRaw.([]interface{})
	if !ok {
		fmt.Printf("failed to typecast to []interface")
		return
	}
	// выводим login и roles
	fmt.Println(login)
	fmt.Println(roles)
}
