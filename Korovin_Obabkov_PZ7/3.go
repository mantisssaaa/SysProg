package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	user := User{
		Username: "John",
		Email:    "John@gmail.com",
	}

	user.SetPassword("1337")
	fmt.Println("Пароль = 1337")
	fmt.Println("Вводим пароль 1337:", user.VerifyPassword("1337"))
	fmt.Println("Вводим неверный пароль:", user.VerifyPassword("1491"))

}

type User struct {
	Username string
	Email    string
	Password []byte
}

func (u *User) SetPassword(password string) {
	hash := sha256.New()
	hash.Write([]byte(password))
	u.Password = hash.Sum(nil)
}

func (u *User) VerifyPassword(password string) bool {
	hash := sha256.New()
	hash.Write([]byte(password))
	return string(hash.Sum(nil)) == string(u.Password)
}
