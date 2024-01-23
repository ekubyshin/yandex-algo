package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
Рита решила оставить у себя одежду только трёх цветов: розового, жёлтого и малинового. После того как вещи других расцветок были убраны, Рита захотела отсортировать свой новый гардероб по цветам. Сначала должны идти вещи розового цвета, потом —– жёлтого, и в конце —– малинового. Помогите Рите справиться с этой задачей.

Примечание: попробуйте решить задачу за один проход по массиву!

Формат ввода
В первой строке задано количество предметов в гардеробе: n –— оно не превосходит 1000000. Во второй строке даётся массив, в котором указан цвет для каждого предмета. Розовый цвет обозначен 0, жёлтый —– 1, малиновый –— 2.

Формат вывода
Нужно вывести в строку через пробел цвета предметов в правильном порядке.

Пример 1
Ввод	Вывод
7
0 2 1 2 0 0 1
0 0 0 1 1 2 2
Пример 2
Ввод	Вывод
5
2 1 2 0 1
0 1 1 2 2
Пример 3
Ввод	Вывод
6
2 1 1 2 0 2
0 1 1 2 2 2
*/

func main() {
	var num int
	fmt.Scanf("%d", &num)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	arr := make([]int, 0, num)
	for i := 0; i < num; i++ {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		arr = append(arr, n)
	}
	res := solve(arr)
	writer := bufio.NewWriter(os.Stdout)
	for i, n := range res {
		writer.WriteString(strconv.Itoa(n))
		if i != num-1 {
			writer.WriteString(" ")
		}
	}
	writer.Flush()
}

func solve(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	counted := make([]int, 3)
	res := make([]int, 0, len(arr))
	for _, n := range arr {
		counted[n] += 1
	}
	for i, n := range counted {
		for j := 0; j < n; j++ {
			res = append(res, i)
		}
	}
	return res
}
