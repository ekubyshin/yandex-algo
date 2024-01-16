package main

import (
	"fmt"
)

/*
Напишите программу, которая определяет, будет ли положительное целое число степенью четвёрки.

Подсказка: степенью четвёрки будут все числа вида 4n, где n – целое неотрицательное число.

Формат ввода
На вход подаётся целое число в диапазоне от 1 до 10000.

Формат вывода
Выведите «True», если число является степенью четырёх, «False» –— в обратном случае.

Пример 1
Ввод	Вывод
15
False
Пример 2
Ввод	Вывод
16
True
*/

func main() {
	var n int
	fmt.Scanln(&n)
	const num = 4
	cur := num
	if n == 1 || n == num {
		fmt.Println("True")
		return
	}
	if n%2 > 0 || n < num {
		fmt.Println("False")
		return
	}
	for {
		if cur >= n {
			break
		}
		cur *= num
	}
	if cur == n {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
