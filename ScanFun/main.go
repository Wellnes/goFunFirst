package main

import "fmt"

func main() {

	for i, s := 0, ""; i < 3; i++ {

		_, резЧтения := fmt.Scan(&s)
		fmt.Println("Ошибки чтения: ", резЧтения)
		fmt.Println("Считанное значение ", s)

	}

}
