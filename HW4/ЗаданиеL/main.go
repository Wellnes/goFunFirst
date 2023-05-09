package main

import "fmt"

func main() {

	fmt.Println("Укажите размерность ряда:")
	var размерностьРяда, знакоСуммаРяда, iСумма int
	fmt.Scan(&размерностьРяда)

	for i := 0; i < размерностьРяда; i++ {

		fmt.Println("Укажите число в ряду:")
		fmt.Scan(&iСумма)

		if i%2 == 1 {
			знакоСуммаРяда -= iСумма
		} else {
			знакоСуммаРяда += iСумма
		}
	}

	fmt.Println(знакоСуммаРяда)

}
