package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var пароль1, пароль2 string

	for паролиДопустимы := false; !паролиДопустимы; {

		fmt.Println("Введите пару паролей:")
		fmt.Scan(&пароль1, &пароль2)

		if ЭтоКороткийПароль(пароль1) {

			fmt.Println("Слишком короткий пароль!")

		} else if ЭтоПростойПароль(пароль1) {

			fmt.Println("Слишком простой пароль!")

		} else if !ПаролиРавны(пароль1, пароль2) {

			fmt.Println("Введенные пароли различаются!")

		} else {

			паролиДопустимы = true
			fmt.Println("Ну наконец-то!")

		}

	}

}

func ЭтоКороткийПароль(пароль string) bool {

	минимальнаяДлинаПароля := 8
	return utf8.RuneCountInString(пароль) < минимальнаяДлинаПароля

}

func ЭтоПростойПароль(пароль string) bool {
	return strings.Contains(пароль, "123") || strings.Contains(пароль, "qwe")
}

func ПаролиРавны(пароль1, пароль2 string) bool {
	return strings.Compare(пароль1, пароль2) == 0
}
