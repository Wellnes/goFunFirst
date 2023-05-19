package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var вхСтрока string

	буферВвода := bufio.NewScanner(os.Stdin)

	захватУдачен := буферВвода.Scan()
	if захватУдачен {
		вхСтрока = буферВвода.Text()
		fmt.Println(вхСтрока)
	}

}
