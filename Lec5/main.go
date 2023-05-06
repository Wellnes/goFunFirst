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

	fmt.Println("===================")

	var val float64 = 0.00
	for i := 0; i < 10; i++ {
		val += 0.10
		fmt.Println(val)
	}
	fmt.Println(val == 1.00)

	fmt.Println("===================")

	fmt.Println(530 % 100)

}
