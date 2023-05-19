package main

import (
	"fmt"
)

func main() {

	имя := "Привет Мир"
	fmt.Println(имя)

	слово := "Слово для примера"
	//словоASRune := []rune(слово)
	fmt.Printf("Строка %s\n", слово)

	for i := 0; i < len(слово); i++ {
		fmt.Printf("%x ", слово[i])
	}

	fmt.Println()

	словоИзСреза := []byte{0xd0, 0xa1, 0xd0, 0xbb, 0xd0, 0xbe, 0xd0, 0xb2, 0xd0, 0xbe, 0x20, 0xd0, 0xb4, 0xd0, 0xbb, 0xd1, 0x8f, 0x20, 0xd0, 0xbf, 0xd1, 0x80, 0xd0, 0xb8, 0xd0, 0xbc, 0xd0, 0xb5, 0xd1, 0x80, 0xd0, 0xb0}
	fmt.Printf("%s\n", словоИзСреза)

	for id, runeVal := range слово {
		fmt.Printf("%c начинается в позиции %d\n", runeVal, id)
	}

	fmt.Println("=======================")

	мойБайтовыйСрез := []byte{0x41, 0x52, 0x41}
	fmt.Println(string(мойБайтовыйСрез))

	мойДесятиричныйБайтовыйСрез := []byte{65, 82, 65}
	fmt.Println(string(мойДесятиричныйБайтовыйСрез))
}
