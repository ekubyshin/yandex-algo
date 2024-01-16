package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
Вечером ребята решили поиграть в игру «Большое число».
Даны числа. Нужно определить, какое самое большое число можно из них составить.

Формат ввода
В первой строке записано n — количество чисел. Оно не превосходит 100.
Во второй строке через пробел записаны n неотрицательных чисел, каждое из которых не превосходит 1000.

Формат вывода
Нужно вывести самое большое число, которое можно составить из данных чисел.

Пример 1
Ввод	Вывод
3
15 56 2
56215
Пример 2
Ввод	Вывод
3
1 783 2
78321
Пример 3
Ввод	Вывод
5
2 4 5 2 10
542210
*/

func main() {
	var n int
	fmt.Scanf("%d", &n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	nums := make([]string, 0, n)
	for i := 0; i < n; i++ {
		if !scanner.Scan() {
			break
		}
		nums = append(nums, scanner.Text())
	}
	fmt.Println(solver(nums))
}

func solver(nums []string) string {
	slices.SortFunc(nums, comparator)
	buf := strings.Builder{}
	for i := len(nums) - 1; i >= 0; i-- {
		buf.WriteString(nums[i])
	}
	return buf.String()
}

func comparator(a string, b string) int {
	if len(a) == len(b) {
		return strings.Compare(a, b)
	}
	return strings.Compare(a+b, b+a)
}
