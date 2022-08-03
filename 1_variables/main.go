// Комментарий для пакета
package main

/*

Многострочный комментарий

*/

func main() {
	var i int = 25    // Битность: int8, int16, int 32, int 64
	var autoInt = -22 //автоматическое определение типа
	var bigInt int64 = 100500
	var unsignedInt uint = 9260          // содержит только положительный диапазон чисел
	var unsignedBigInt uint64 = 10005000 //  Битность: uint8, uint16, uint 32, uint 64
	println("integers: ", i, autoInt, bigInt, unsignedInt, unsignedBigInt)

	// return

	//Тип с плавающей точкой
	//Битность float32 float64
	var f32 float32 = 3.134
	var f64 float64 = 3.1234
	println("float: ", f32, f64)

	// Булевы значения
	var b bool = true
	println("flag is: ", b)

	//Строки
	var hello string = "Hi!\n\t" // спецсимволы \n - новая строка, \t - табуляция
	var world = "world"          //автоматическое присваивание типа
	println(hello, world)

	//Переменные имеют значения по умолчанию
	var defaultInt int
	var defaultFloat float32
	var defaultString string
	var defaultBool bool
	println(defaultInt, defaultFloat, defaultString, defaultBool)

	//Тип byte хранит в себе символьное значение
	//'\x25' представление символа % в шестнадцатиричной системе счичления из таблицы ASCII
	// Alias для uint8 - [0;255]
	var rawSymbol byte = '\x25' // или 37 код в 10-й системе счисления
	println("symbol:", string(rawSymbol))

	//Краткое объявление переменной
	shortDeclare := 42 //
	println("Short declare:", shortDeclare)

	//Приведение типов
	println("float to int:", int(f32))
	// Для приведения чисел в string нужен пакет strconv для работы с типизацияей цифр и строк
	println("int to string:", string(51)) // ASCII, DEC 51 - Chr 3

	//комплексные числа
	z := 2 + 3i
	println("complex number", z)

	//операции со строками
	s1 := "Nikita"
	s2 := "Borozdin"
	fullName := s1 + s2
	println("My name is: ", fullName, len(fullName))

	//Многострочное объявление
	escaping := `Hello\r\n
	World`
	println("<pre>: ", escaping, "</pre>")

	//Объявление нескольких переменных
	var v1, v2 string = "v1", "v2"
	println(v1, v2)

	//Объявление группы переменных
	var (
		m0 int = 12
		m2     = "string"
		m3     = 23
	)

	println(m0, m2, m3)
}
