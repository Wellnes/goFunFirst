package main

import "fmt"

func main() {

	mapper := make(map[string]int)
	fmt.Printf("%T\n", mapper)

	mapper["Alice"] = 24
	mapper["Bob"] = 25

	fmt.Println(mapper)

	newMapper := map[string]int{
		"Bob":   1000,
		"Alice": 1000,
	}
	newMapper["Jo"] = 1000

	fmt.Println(newMapper)

	//	var mapZeroValue map[string]int
	//	mapZeroValue["Alice"] = 34			// panic: assignment to entry in nil map

	знач, ок := newMapper["ARA"]
	fmt.Println(знач, ок)

	for ключ, знч := range newMapper {
		fmt.Println(ключ, знч)
	}

}
