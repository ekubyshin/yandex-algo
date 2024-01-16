package main

import (
	"fmt"
	"strings"
)

/*
На клавиатуре старых мобильных телефонов каждой цифре соответствовало несколько букв. Примерно так:

2:'abc',
3:'def',
4:'ghi',
5:'jkl',
6:'mno',
7:'pqrs',
8:'tuv',
9:'wxyz'

Вам известно в каком порядке были нажаты кнопки телефона, без учета повторов. Напечатайте все комбинации букв, которые можно набрать такой последовательностью нажатий.

Формат ввода
На вход подается строка, состоящая из цифр 2-9 включительно. Длина строки не превосходит 10 символов.

Формат вывода
Выведите все возможные комбинации букв через пробел в лексикографическом (алфавитном) порядке по возрастанию.
1.
23 ad ae af bd be bf cd ce cf
2.
92 wa wb wc xa xb xc ya yb yc za zb zc
*/

var symbols = map[byte][]string{
	'2': {"a", "b", "c"},
	'3': {"d", "e", "f"},
	'4': {"g", "h", "i"},
	'5': {"j", "k", "l"},
	'6': {"m", "n", "o"},
	'7': {"p", "q", "r", "s"},
	'8': {"t", "u", "v"},
	'9': {"w", "x", "y", "z"},
}

func main() {
	var str string
	fmt.Scanln(&str)
	fmt.Println(strings.Join(generate(str), " "))
}

func generate(s string) []string {
	if len(s) == 0 {
		return []string{""}
	}
	res := make([]string, 0, len(s)*4)
	tail := generate(s[1:])
	for _, c := range symbols[s[0]] {
		for _, sc := range tail {
			res = append(res, c+sc)
		}
	}
	return res
}
