package main

import "fmt"

var globalMap map[string]string

//Функция init нужна для работы до выполнения функции main
func init() {
	globalMap = make(map[string]string) // инициализируем
	globalMap["myName"] = "Nick"
}

func main() {
	var mm map[string]string
	fmt.Println("Объявленная, но не инициализированная карта", mm)
	// panic сигнализирует о наличии неполадок, из-за которых система не может продолжать функционировать
	// mm["test"] = "test" - panic

	// полная инициализация
	mm2 := map[string]string{} //или var mm2 map[int]int = map[int]int{}
	mm2["test"] = "ok"
	fmt.Println(mm2)

	// короткая инициализация
	var mm3 = make(map[string]string)
	mm3["myName"] = "Nikita"
	fmt.Println(mm3)

	fmt.Println("globalMap", globalMap)
}
