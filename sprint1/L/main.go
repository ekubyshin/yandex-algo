package main

import (
	"fmt"
)

/*
Васе очень нравятся задачи про строки, поэтому он придумал свою. Есть 2 строки s и t, состоящие только из строчных букв. Строка t получена перемешиванием букв строки s и добавлением 1 буквы в случайную позицию. Нужно найти добавленную букву.

Формат ввода
На вход подаются строки s и t, разделённые переносом строки. Длины строк не превосходят 1000 символов. Строки не бывают пустыми.

Формат вывода
Выведите лишнюю букву.

Пример 1
Ввод	Вывод
abcd
abcde
e
Пример 2
Ввод	Вывод
go
ogg
g
Пример 3
Ввод	Вывод
xtkpx
xkctpx
*/

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	am := make(map[rune]int)
	var unknown rune
	for _, c := range a {
		if v, ok := am[c]; ok {
			am[c] = v + 1
		} else {
			am[c] = 1
		}
	}
	for _, c := range b {
		if v, ok := am[c]; ok {
			if v == 0 {
				unknown = c
				break
			}
			am[c] -= 1
		} else {
			unknown = c
			break
		}
	}
	fmt.Println(string(unknown))
}
