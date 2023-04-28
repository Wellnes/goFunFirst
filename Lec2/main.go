package main

import (
	"fmt"
	"math"
)

func main() {

	// простейший вывод на консоль
	fmt.Println("Hello world")
	fmt.Println("Second", "line")

	// Функция Print - простой вывод аргумента
	fmt.Print("First")
	fmt.Print("Second", 42, 13, "test")
	fmt.Print("Third\n")

	// Форматированный вывод: Printf - стандартный вывод с флагами форматирования
	fmt.Printf("Hello, my name is %s\nMy age is %d\n", "Bob", 42)

	Сообщить("Привет медвед")

	// --------------------------

	var возраст int
	возраст = 32
	fmt.Printf("Возраст %d \n", возраст)
	fmt.Println("Мой возвраст:", возраст)

	var мойВес int = 183
	fmt.Println("Мой вес:", мойВес)

	var мойРост = 175
	fmt.Println("Мой рост:", мойРост)

	счетчик := 10
	fmt.Println("Счетчик: ", счетчик)

	счетчик++
	fmt.Println("Счетчик: ", счетчик)

	счетчик += 1
	fmt.Println("Счетчик: ", счетчик)

	// Вывести тип переменной
	fmt.Printf("%T\n", счетчик)

	fmt.Println("(^_^)")

	aArg, bArg := 10, "Vova"
	fmt.Println(aArg, bArg)

	ширина, высота := 30.2, 23.5
	fmt.Printf("Минимальная сторона прямоугольника составляет: %.2f\n", math.Min(ширина, высота))

	// практика

	var (
		слово1 = "имя"
		слово2 = "твое"
		слово3 = "мне"
		слово4 = "знакомо"
	)
	fmt.Println(слово4, слово3, слово2, слово1)
	fmt.Println(слово3, слово1, слово4, слово2)
	fmt.Println(слово2, слово4, слово1, слово3)

}

func Сообщить(ВыхСтрока string) {

	fmt.Println(ВыхСтрока)

}
