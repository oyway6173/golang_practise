package main

import "fmt"

func main() {
	var sl []int
	fmt.Println("пустой слайс: ", sl)

	sl = append(sl, 100) //Добаление элемента в слайс
	fmt.Println("Слайс с 1-м элементом ", sl)
	//Пример использования: хранение номеров пользователей на удаление
	// Отличие слайсов от массива - слайсы изменить можно, массивы - нет
}
