package main

import "fmt"

func main() {

	var (
		возраст int
		имя     string
	)

	fmt.Println("Введите ваш возраст и имя:")
	fmt.Scan(&возраст, &имя)
	fmt.Printf("Мой возраст: %d\nМоё имя: %s\n", возраст, имя)

}
