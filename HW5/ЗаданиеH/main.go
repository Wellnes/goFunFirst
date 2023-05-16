package main

import "fmt"

func main() {

	var считаннаяСтрока string
	// rez, _ := fmt.Scan(&считаннаяСтрока)
	// fmt.Println(rez, считаннаяСтрока)

	// rez, _ = fmt.Scan(&считаннаяСтрока)
	// fmt.Println(rez, считаннаяСтрока)

	// rez, _ = fmt.Scan(&считаннаяСтрока)
	// fmt.Println(rez, считаннаяСтрока)

	for rez, _ := fmt.Scan(&считаннаяСтрока); rez == 1; {
		fmt.Println(rez, считаннаяСтрока)
	}

}
