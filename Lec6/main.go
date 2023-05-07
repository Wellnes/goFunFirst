package main

import (
	"fmt"
	"strings"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("==========")
	tempS := "Hello, Kitty!"
	for _, буква := range tempS {
		fmt.Printf("%c\n", буква)
	}
	fmt.Println("==========")

	// Елочка вложенным циклом
	for i := 0; i < 10; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
	}

	//
	var пароль string

	for {
		fmt.Println("Укажите пароль:")
		fmt.Scan(&пароль)

		if strings.Contains(пароль, "123") {
			fmt.Println("Пароль слишком слабый, попробуйте еще раз")
		} else {
			fmt.Println("Пароль принят")
			break
		}

	}

	// Цикл с множественными переменными цикла
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)
	}

}
