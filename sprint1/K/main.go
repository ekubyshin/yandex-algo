package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Вася просил Аллу помочь решить задачу. На этот раз по информатике.

Для неотрицательного целого числа X списочная форма –— это массив его цифр слева направо. К примеру, для 1231 списочная форма будет [1,2,3,1]. На вход подается количество цифр числа Х, списочная форма неотрицательного числа Х и неотрицательное число K. Число К не превосходят 10000. Длина числа Х не превосходит 1000.

Нужно вернуть списочную форму числа X + K.

Не используйте встроенные средства языка для сложения длинных чисел.

Формат ввода
В первой строке — длина списочной формы числа X. На следующей строке — сама списочная форма с цифрами записанными через пробел.

В последней строке записано число K, 0 ≤ K ≤ 10000.

Формат вывода
Выведите списочную форму числа X+K.

Пример 1
Ввод	Вывод
4
1 2 0 0
34
1 2 3 4
Пример 2
Ввод	Вывод
2
9 5
17
*/

func main() {
	var n int
	fmt.Scanln(&n)
	arr := make([]string, n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arr[i] = scanner.Text()
	}
	scanner.Scan()
	s := scanner.Text()
	b := strings.Split(s, "")
	maxLen := len(arr)
	minLen := len(b)
	maxN := &arr
	minN := &b
	if maxLen < len(b) {
		maxLen = len(b)
		minLen = len(arr)
		maxN = &b
		minN = &arr
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
		if sum >= 10 {
			p = 1
			sum = sum - 10
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
		if i != len(buf)-1 {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
}
