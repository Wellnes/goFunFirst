package main

import "fmt"

func main() {

	startArr := [4]int{10, 12, 34, 56}
	var startSlice = startArr[1:4]
	startSlice = append(startSlice, 45)
	startSlice = append(startSlice, 63)
	fmt.Println("Срез", startSlice, "Длина среза:", len(startSlice), "Ёмкость", cap(startSlice))
	fmt.Println("Массив", startArr)

	startSlice = append(startSlice, 72)
	fmt.Println("Срез", startSlice, "Длина среза:", len(startSlice), "Ёмкость", cap(startSlice))
	startSlice = append(startSlice, 73)
	fmt.Println("Срез", startSlice, "Длина среза:", len(startSlice), "Ёмкость", cap(startSlice))
	startSlice = append(startSlice, 74)
	fmt.Println("Срез", startSlice, "Длина среза:", len(startSlice), "Ёмкость", cap(startSlice))

	fmt.Println("=================")
	sl := make([]int, 10, 15)
	fmt.Println(sl)

	fmt.Println("=================")
	моиСлова := []string{"один", "два", "три"}
	fmt.Println(моиСлова)
	моиСлова = append(моиСлова, "четыре")
	fmt.Println(моиСлова)

	другойСлайз := []string{"пять", "шесть"}
	резСлайз := append(моиСлова, другойСлайз...)
	fmt.Println(резСлайз)

	fmt.Println("=================")

	мСрез := [][]int{
		{1, 2},
		{3, 4, 5, 6},
		{10, 20, 30},
		{},
	}

	fmt.Println(мСрез)

	fmt.Println("=================")

	байтовыйСрез := []byte{97, 36, 68, 73, 77, 65, 32, 196, 200, 204, 192}

	for _, val := range байтовыйСрез {
		fmt.Printf("%c\n", val)
	}
}
