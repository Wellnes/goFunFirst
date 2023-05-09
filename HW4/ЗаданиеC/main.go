package main

import "fmt"

func main() {

	var число int32

	for {

		fmt.Scan(&число)
		if число == 0 {
			return
		} else {
			fmt.Println(число)
		}

	}

}
