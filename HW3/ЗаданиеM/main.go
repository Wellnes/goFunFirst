package main

import (
	"fmt"
	"strings"
)

func main() {
	var слово1, слово2, слово3 string
	fmt.Println("Произведите расчет")

	fmt.Scan(&слово1, &слово2, &слово3)

	слово1Исправно := strings.Compare(слово1, "раз") == 0 || strings.Compare(слово1, "один") == 0
	слово2Исправно := strings.Compare(слово2, "два") == 0
	слово3Исправно := strings.Compare(слово3, "три") == 0

	if слово1Исправно && слово2Исправно && слово3Исправно {
		fmt.Println("ОК")
	} else if strings.Compare(слово1, "1") == 0 && strings.Compare(слово2, "2") == 0 && strings.Compare(слово3, "3") == 0 {
		fmt.Println("ОК")
	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}

}
