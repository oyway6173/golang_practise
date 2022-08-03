package main

//Глобальные константы - некие величины, не изменяющие свои значения
const myInt int32 = 17
const someInt = 1
const myName = "Nikita"

//Группа констант
const (
	flag1 = 1
	flag2 = 2
)

//Группа констант + iota - генератор инкримента
//Eнам
const (
	zeroVar = iota
	oneVar
	_        //пустая переменная, пропуск iota
	threeVar // 3
)

//Арифметические действия
const (
	z1 = iota + 1
	z2
	z3
)

func main() {
	println(zeroVar, oneVar, threeVar)
}
