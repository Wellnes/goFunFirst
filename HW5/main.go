package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	var квоСериаловСервиса, квоИнтересующихСериалов int

	fmt.Println("Введите колво сериалов и их названия (подробней см. справку)")
	fmt.Scan(&квоСериаловСервиса, &квоИнтересующихСериалов)

	мсИнтересующиеСериалы := [...]string{5: ""}
	мсСериалыСервиса := [...]string{5: ""}

	input := bufio.NewScanner(os.Stdin)
	input.Scan()

	for i := 0; i < квоСериаловСервиса; i++ {

		input.Scan()
		мсСериалыСервиса[i] = input.Text()

	}

	for i := 0; i < квоИнтересующихСериалов; i++ {

		input.Scan()
		мсИнтересующиеСериалы[i] = input.Text()

	}

	// fmt.Printf("%q\n", мсСериалыСервиса)
	// fmt.Printf("%q\n", мсИнтересующиеСериалы)

	for id := 0; id < квоИнтересующихСериалов; id++ {

		if strings.Compare(мсИнтересующиеСериалы[id], мсСериалыСервиса[id]) == 0 {
			fmt.Println("ЕСТЬ")
		} else {
			fmt.Println("НЕТ В НАЛИЧИИ")
		}

	}

}
