package main

import (
	"fmt"
	"user"

	lvl "github.com/google/level"
)

/*
Глобальные переменные:
GOPATH - путь к проекту
GOROOT - путь к месту, где установили GO

Структура проекта:
	bin/ - тут лежат бинарники
	pkg/ - тут лежат объектные файлы, корторые создаются при компиляции
	src/ - тут лежит ваш код и другие пакеты

Exporting:
	Если переменная, константа, функция, структура начинается с маленькой буквы, то мы ее не можем использовать в других пакетах, только в текущем
	Если переменная, константа, функция, структура с большой буквы, то мы можем использовать ее в других пакетах

*/

func main() {
	fmt.Println("ADMIN TYPE", user.Username)
	fmt.Println("Type admin", user.TypeAdmin)
	fmt.Println("Type user", user.TypeUser)
	fmt.Println("Level", lvl.LevelOne)
}
