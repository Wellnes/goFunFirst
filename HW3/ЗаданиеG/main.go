package main

import "fmt"

func main() {

	var aStorona, bStorona int

	fmt.Println("ВВедите длины сторон предполагаемого прямоугольника:")

	fmt.Scan(&aStorona, &bStorona)

	fmt.Println("Периметр:", 2*(aStorona+bStorona))
	fmt.Println("Площадь:", aStorona*bStorona)

}
