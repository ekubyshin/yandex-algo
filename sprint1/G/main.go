package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Вася реализовал функцию, которая переводит целое число из десятичной системы в двоичную. Но, кажется, она получилась не очень оптимальной.

Попробуйте написать более эффективную программу.

Не используйте встроенные средства языка по переводу чисел в бинарное представление.

Формат ввода
На вход подаётся целое число в диапазоне от 0 до 10000.

Формат вывода
Выведите двоичное представление этого числа.

Пример 1
Ввод	Вывод
5
101
Пример 2
Ввод	Вывод
14
1110
*/

func main() {
	var n int
	fmt.Scanln(&n)
	res := make([]int, 0, 10)
	cur := n
	for cur > 1 {
		rem := cur % 2
		cur = (cur - rem) / 2
		res = append(res, rem)
	}
	res = append(res, cur)
	writer := bufio.NewWriter(os.Stdout)
	for i := len(res) - 1; i >= 0; i-- {
		output_string := strconv.Itoa(res[i])
		writer.WriteString(output_string)
	}
	writer.Flush()
}
