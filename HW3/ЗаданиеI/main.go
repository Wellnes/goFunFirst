package main

import "fmt"

func main() {

	fmt.Println("Введите три натуральных числа")

	var s1, s2, s3 int32
	fmt.Scan(&s1, &s2, &s3)

	fmt.Printf("%d%d%d\n", s3, s2, s1)
}
