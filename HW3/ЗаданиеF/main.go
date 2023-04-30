package main

import "fmt"

func main() {

	var строка1, строка2, строка3, строка4, имяАвтора string
	fmt.Scan(&строка1, &строка2, &строка3, &строка4, &имяАвтора)

	fmt.Println(строка4, "-", имяАвтора)
	fmt.Println(строка3, "-", имяАвтора)
	fmt.Println(строка2, "-", имяАвтора)
	fmt.Println(строка1, "-", имяАвтора)

}
