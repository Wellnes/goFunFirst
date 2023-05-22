package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	files := os.Args[1:2]
	firstFileName := files[0]
	f, err := os.Open(firstFileName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "файл: %v\n", err)
		return
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		fmt.Println(input.Text())
	}

	f.Close()
}
