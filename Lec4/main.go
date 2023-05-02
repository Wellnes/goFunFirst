package main

import (
	"fmt"
	"os"
	"unsafe"
)

var globalHawk string = "ARA"

func main() {

	var этоПокемон bool
	fmt.Println(этоПокемон)

	aBoolean, bBoolean := true, false
	fmt.Println(aBoolean && bBoolean)
	fmt.Println(aBoolean || bBoolean)
	fmt.Println(!aBoolean)

	//

	var a, b int8 = 32, 92

	fmt.Println("Значение а:", a, "Значение b:", b, "a+b:", a+b)

	// узаем разрядность

	var aInt = 32
	fmt.Printf("Разрядность инта: %d байт\n", unsafe.Sizeof(aInt))

	var first32 int32 = 123
	var second64 int64 = 678

	fmt.Println(int64(first32) + second64)

	// вещественные числа
	var floatFirst, floatSecond float32 = 5.67, 12.64
	fmt.Printf("Тип первого числа это %T, тип второго числа это %T\n", floatFirst, floatSecond)

	// комплексные числа
	c1 := complex(5, 7)
	c2 := 12 + 32i
	println(c1 + c2)

	fmt.Println(globalHawk)

	ВывестиРазделитель()
	fmt.Println(os.Args[0])

	ВывестиРазделитель()
	var runeTest rune = 'D'
	fmt.Printf("Char тип: %T\n", runeTest)
	fmt.Printf("Char тип: %c\n", runeTest)
	fmt.Printf("Char тип: %d\n", runeTest)

}

func ВывестиРазделитель() {
	fmt.Println("---------------------------------------")
}
