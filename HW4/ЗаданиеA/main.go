package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {

	fmt.Println("Введите текстовый логотип:")

	var текстовыйЛоготип string
	fmt.Scan(&текстовыйЛоготип)

	числоБукв := utf8.RuneCountInString(текстовыйЛоготип)
	fixPrice := 0.23

	суммаЗаказа := float64(числоБукв) * fixPrice

	if суммаЗаказа < 1 {
		fmt.Printf("%.f коп.\n", суммаЗаказа*100)
	} else {
		рубли := math.Floor(суммаЗаказа)
		копейки := суммаЗаказа - рубли
		fmt.Printf("%.f р. %.f коп. \n", рубли, копейки*100)
	}

}
