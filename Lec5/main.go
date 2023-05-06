package main

import (
	"fmt"
	"strings"
)

func main() {

	var value int32

	fmt.Println("Введите значение:")
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("Значение четное")
	} else {
		fmt.Println("Значение не четное")
	}

	if isDno := false; isDno {
		fmt.Println("====Подвал=====")
	}

	fmt.Println("Введите цвет")
	var color string
	fmt.Scan(&color)

	if color == "green" {

		fmt.Println("Цвет зеленый")

	} else if strings.Compare(color, "red") == 0 {

		fmt.Println("Цвет красный")

	} else {
		fmt.Println("Обнаружен не ожидаемый цвет")
	}

	// Инициализация в блоке условного оператора

	if num := 10; num%2 == 0 {
		fmt.Println("num Четный")
	} else {
		fmt.Println("num НЕ Четный")
	}

}
