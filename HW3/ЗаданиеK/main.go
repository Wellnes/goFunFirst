package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Введите комментарий:")

	var комментарий string
	fmt.Scan(&комментарий)

	if strings.Contains(комментарий, "чат") {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}

}
