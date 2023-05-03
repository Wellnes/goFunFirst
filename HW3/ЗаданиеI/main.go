package main

import "fmt"

func main() {

	fmt.Println("Введите три натуральных числа")

	var s1, s2, s3 string
	fmt.Scan(&s1, &s2, &s3)

	fmt.Print(s3, s2, s1, "\n")
}
