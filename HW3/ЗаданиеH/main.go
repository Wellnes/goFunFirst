package main

import (
	"fmt"
	"math"
)

func main() {
	var число1, число2 float64

	fmt.Println("Введите 2 числа:")
	fmt.Scan(&число1, &число2)
	fmt.Println("Квадрат суммы введеных чисел составляет:", math.Pow(число1+число2, 2))

}
