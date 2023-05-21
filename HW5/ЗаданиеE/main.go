package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var вхСтрока string

	буфВвод := bufio.NewScanner(os.Stdin)

	if буфВвод.Scan() {
		вхСтрока = буфВвод.Text()
	}

	номПозиции := 0
	for _, текБуква := range вхСтрока {

		номПозиции++

		if номПозиции%2 == 0 {
			continue
		}

		for i := 0; i < 3; i++ {
			fmt.Printf("%c", текБуква)
		}

	}

}
