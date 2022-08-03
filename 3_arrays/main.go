package main

import "fmt"

const size uint = 5

func main() {
	var a1 [5]int
	fmt.Println("Массив:", a1, "длина:", len(a1))

	//Устанавливаем размер массива используя константу
	var a2 [5 * size]bool
	fmt.Println("Массив2:", a2, "длина:", len(a2))

	//Устанавливаем длину при инициализации
	a3 := [...]int{1, 2, 3}
	fmt.Println("Массив3:", a3, "длина при инициализации:", len(a3))

	//Двухмерный массив
	var aa [3][3]int
	aa[1][1] = 1
	fmt.Println("двухмерный масссив:", aa)

	return
}
