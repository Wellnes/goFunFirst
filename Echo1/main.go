// Echo1 выводит аргументы командной строки
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	start := time.Now()

	// var s, sep string
	// for i := 1; i < len(os.Args); i++ {
	// 	s += sep + os.Args[i]
	// 	sep = " "
	// }
	// fmt.Println(s)

	fmt.Println(strings.Join(os.Args[1:], " "))

	//	for id, arg := range os.Args {
	//		fmt.Println("Индекс:", id, "Значение:", arg)
	//	}

	miliSecs := time.Since(start).Milliseconds()
	fmt.Println("Прошло милисекунд:", miliSecs)

}
