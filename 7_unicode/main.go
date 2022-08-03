package main

import (
	"fmt"
)

/*
	Юникод - стандарт кодирования символов, включающий в себя знаки почти всех письменных языков мира

	коды в стандарте Юникод разделены на несколько областей
	К примеру, область с кодами от U+0000 до U+007F содержит символы набора ASCII, и коды этих символов совпадают с кодами в таблицe ASCII
*/

func main() {
	var symbol rune = 'Г' // можем указать юникодный символ прямо в коде
	unicodeSymbol := '%'  // можем указать юникодный символ прямо в коде
	unicodeSymbolByNumber := '\u2318'
	fmt.Println(symbol /*порядковый номер*/, fmt.Sprintf("%#U", unicodeSymbol), string(unicodeSymbolByNumber))

	//ASCII symbols
	strl := "BorozdinNikita"
	//выводим строчку и ее длину в байтах
	fmt.Println(strl, len(strl), strl[0])

	//Unicode symbols
	str2 := "БороздинНикита"
	//выводим строчку и ее длину в байтах
	fmt.Println(str2, len(str2))

	//переводим строку в слайс байт, чтобы мы могли проитерироавться по байтам, а не по символам
	//Юникод это многобайтная кодировка, и в данном случае символ 'Б' будет занимать два байта
	bin := []byte(str2)
	fmt.Println("В двоичном виде", bin, len(bin), string(bin[:2]))

	//Итерируемся по байтам
	for i := 0; i < len(bin); i++ {
		val := bin[i]
		fmt.Println(i, val)
	}

	fmt.Println("Range")

	for idx, val := range strl {
		fmt.Println(idx, string(val))
	}
}
