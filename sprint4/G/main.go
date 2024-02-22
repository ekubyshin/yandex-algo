package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Жители Алгосов любят устраивать турниры по спортивному программированию. Все участники разбиваются на пары и соревнуются друг с другом. А потом два самых сильных программиста встречаются в финальной схватке, которая состоит из нескольких раундов. Если в очередном раунде выигрывает первый участник, в таблицу с результатами записывается 0, если второй, то 1. Ничьей в раунде быть не может.

Нужно определить наибольший по длине непрерывный отрезок раундов, по результатам которого суммарно получается ничья. Например, если дана последовательность 0 0 1 0 1 1 1 0 0 0, то раунды с 2-го по 9-й (нумерация начинается с единицы) дают ничью.

Формат ввода
В первой строке задаётся n (0 ≤ n ≤ 105) –— количество раундов. Во второй строке через пробел записано n чисел –— результаты раундов. Каждое число равно либо 0, либо 1.

Формат вывода
Выведите длину найденного отрезка.
*/

func main() {
	sequence := read()
	fmt.Println(solve(sequence))
}

func solve(sequence []int) int {
	balances := map[int]int{0: -1}
	balance := 0
	maxLength := 0
	for i, s := range sequence {
		if s == 1 {
			balance++
		} else {
			balance--
		}
		if _, ok := balances[balance]; !ok {
			balances[balance] = i
			continue
		}
		length := i - balances[balance]
		if length > maxLength {
			maxLength = length
		}
	}
	return maxLength
}

func read() []int {
	data, _ := os.ReadFile("./input.txt")
	dataArray := strings.Split(string(data), "\n")
	rounds, _ := strconv.Atoi(dataArray[0])
	sequence := make([]int, 0, rounds)
	for _, line := range strings.Split(dataArray[1], " ") {
		result, _ := strconv.Atoi(line)
		sequence = append(sequence, result)
	}
	return sequence
}
