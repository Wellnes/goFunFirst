package main

import "fmt"

func main() {

	// var price int
	// fmt.Scan(&price)

	// switch price {
	// case 100:
	// 	fmt.Println("Первый случай")
	// case 110:
	// 	fmt.Println("Второй случай")
	// case 120:
	// 	fmt.Println("Третий случай")
	// default:
	// 	fmt.Println("Cлучай по умолчанию")
	// }

	// var возрастнаяГруппа rune = 'b'

	// switch возрастнаяГруппа {
	// case 'a', 'b', 'c':
	// 	fmt.Println("Возрастная группа 10 - 30")
	// case 'd', 'e':
	// 	fmt.Println("Возрастная группа 40 - 50")
	// default:
	// 	fmt.Println("Вы слишком молоды или стары")
	// }

	var number int
	fmt.Scan(&number)

	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		fallthrough // как бы отменяет вшитый break и передает управление следующему кейсу, даже не проверяя условие
	case number > 200:
		fmt.Printf("%d is less then 200\n", number)
	}

}
