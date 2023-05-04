package main

import "fmt"

func main() {

	fmt.Println("Введите два вещественных числа")

	var число1, число2 float32
	fmt.Scan(&число1, &число2)

	суммаЧисел := число1 + число2
	if int32(суммаЧисел)%2 == 0 {
		fmt.Println("ЧЁТНОЕ")
	} else {
		fmt.Println("НЕЧЁТНОЕ")
	}

}
