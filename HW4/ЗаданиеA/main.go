package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Введите текстовый логотип:")

	var текстовыйЛоготип string
	fmt.Scan(&текстовыйЛоготип)

	числоБуквЛоготипа := utf8.RuneCountInString(текстовыйЛоготип)
	fixPriceКопеек := 23

	суммаЗаказаКопеек := числоБуквЛоготипа * fixPriceКопеек
	одинРубль := 1 * CentFactor()

	if суммаЗаказаКопеек < одинРубль {
		fmt.Printf("%.d коп.\n", суммаЗаказаКопеек)
	} else {
		остаток := суммаЗаказаКопеек % CentFactor()
		целое := (суммаЗаказаКопеек - остаток) / CentFactor()
		fmt.Printf("%d р. %d коп.\n", целое, остаток)
	}

}

func CentFactor() int {
	return 100
}
