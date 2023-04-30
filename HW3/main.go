package main

import "fmt"

func main() {

	// ЗаданиеD()
	ЗаданиеE()

}

func ЗаданиеE() {

	var (
		строка1 string
		строка2 string
		строка3 string
		строка4 string
	)

	fmt.Println("Введите четыре строки:")
	fmt.Scan(&строка1, &строка2, &строка3, &строка4)
	fmt.Printf("I wanna some '%s', my favorite meal - '%s'. Give me also '%s'. I will come soon... '%s'", строка1, строка2, строка3, строка4)

}

func ЗаданиеD() {

	var (
		имя     string
		фамилия string
		возраст int
	)

	fmt.Println("Введите имя, фамилию и возраст:")
	fmt.Scan(&имя, &фамилия, &возраст)
	fmt.Printf("Имя: %s , Фамилия: %s , Возраст: %d . Студент BPS", имя, фамилия, возраст)

}
