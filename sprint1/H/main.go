package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Тимофей записал два числа в двоичной системе счисления и попросил Гошу вывести их сумму, также в двоичной системе. Встроенную в язык программирования возможность сложения двоичных чисел применять нельзя. Помогите Гоше решить задачу.

Решение должно работать за O(N), где N –— количество разрядов максимального числа на входе.

Формат ввода
Два числа в двоичной системе счисления, каждое на отдельной строке. Длина каждого числа не превосходит 10 000 символов.

Формат вывода
Одно число в двоичной системе счисления.

Пример 1
Ввод	Вывод
1010
1011
10101
Пример 2
Ввод	Вывод
1
1
10
*/

func main() {
	var a, b string
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	a = a[:]
	b = b[:]
	maxLen := len(a)
	minLen := len(b)
	maxN := &a
	minN := &b
	if maxLen < len(b) {
		maxLen = len(b)
		minLen = len(a)
		maxN = &b
		minN = &a
	}
	j := minLen - 1
	p := 0
	buf := make([]int, maxLen+1)
	for i := maxLen - 1; i >= 0; i-- {
		ai := 0
		bj := 0
		ai, _ = strconv.Atoi(string((*maxN)[i]))
		if j >= 0 {
			bj, _ = strconv.Atoi(string((*minN)[j]))
			j -= 1
		} else {
			bj = 0
		}
		sum := ai + bj + p
		if sum == 3 {
			sum = 1
			p = 1
		} else if sum == 2 {
			sum = 0
			p = 1
		} else {
			p = 0
		}
		buf[i+1] = sum
	}
	buf[0] = p
	writer := bufio.NewWriter(os.Stdout)
	for i, v := range buf {
		if i == 0 && v == 0 {
			continue
		}
		writer.WriteString(strconv.Itoa(v))
	}
	writer.Flush()
}
