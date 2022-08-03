package main

func main() {
	// a := true
	var a bool = true
	if a {
		println("argument a is true")
	}

	b := 1
	if b == 1 {
		println("только явное преобразование (никаких if b)")
	}

	name := "Nik"
	if len(name) > 5 {
		println("name len > 5", name)
	} else {
		println("name len <= 4")
	}

	if len(name) > 5 {
		println("name len > 5", name)
	} else if len(name) == 4 {
		println("name len = 4", name)
	} else {
		println("name len < 4", name)
	}

	for {
		println("бесконечный цикл")
		break
	}

	i := 0
	for i < 5 {
		if i == 1 {
			i++
			continue //переходим к следующей итерации цикла
		}
		println("while цикл, i:", i)
		i++
	}

	sl := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(sl); i++ {
		println("Обычный цикл", i, sl[i])
	}

	//Может проитероваться по слайду
	for idx, val := range sl {
		println("цикл range", idx, val)
	}

	var switchVar = 2
	switch switchVar {
	case 1:
		println("Кейс 1")
	case 2:
		println("Кейс 2")
		fallthrough //проваливаемся в default case
	default:
		println("Дефолтный кейс")
	}

	//Метки
	index := 1
Loop:
	for {
		index++
		switch index {
		case 10:
			println("прерывание метки Loop")
			break Loop //прерывание метки Loop
		}
	}

LoopOuter:
	for k := 0; k < 10; k++ {
		for j := 0; j < 10; j++ {
			println("k, j =", k, j)
			break LoopOuter // в таком случае мы выходим из всех циклов сразу
		}
	}
}
