package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var login, email string

	fmt.Println("Укажите логин и пароль:")
	fmt.Scan(&login, &email)

	if !КорректныйЛогин(login) {
		fmt.Println("Некорректный логин")
	} else if !КорректнаяПочта(email) {
		fmt.Println("Некорректная почта")
	} else {
		fmt.Println("ОК")
	}

}

func КорректнаяПочта(почта string) bool {

	rez := true

	if !strings.Contains(почта, ".") || !strings.Contains(почта, "@") {
		rez = false
	}

	return rez

}

func КорректныйЛогин(login string) bool {

	rez := true

	if utf8.RuneCountInString(login) < 10 || strings.Contains(login, "@") {
		rez = false
	}

	return rez

}
