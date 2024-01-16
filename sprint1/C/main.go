package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Дана матрица. Нужно написать функцию, которая для элемента возвращает всех его соседей. Соседним считается элемент, находящийся от текущего на одну ячейку влево, вправо, вверх или вниз. Диагональные элементы соседними не считаются.

Например, в матрице A соседними элементами для (0, 0) будут 2 и 0. А для (2, 1) –— 1, 2, 7, 7.

Формат ввода
В первой строке задано n — количество строк матрицы. Во второй — количество столбцов m. Числа m и n не превосходят 1000. В следующих n строках задана матрица. Элементы матрицы — целые числа, по модулю не превосходящие 1000. В последних двух строках записаны координаты элемента, соседей которого нужно найти. Индексация начинается с нуля.

Формат вывода
Напечатайте нужные числа в возрастающем порядке через пробел.

Пример 1
Ввод	Вывод
4
3
1 2 3
0 2 6
7 4 1
2 7 0
3
0
7 7
Пример 2
Ввод	Вывод
4
3
1 2 3
0 2 6
7 4 1
2 7 0
0
0
*/

func main() {
	var n, m, x, y int
	fmt.Scanln(&n)
	fmt.Scanln(&m)
	arr := make([][]int, n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		values := strings.Split(line, " ")
		l := make([]int, m)
		arr[i] = l
		for j := 0; j < m; j++ {
			v, _ := strconv.Atoi(values[j])
			arr[i][j] = v
		}
	}
	scanner.Scan()
	k := scanner.Text()
	y, _ = strconv.Atoi(k)
	scanner.Scan()
	k = scanner.Text()
	x, _ = strconv.Atoi(k)
	res := make([]int, 0, 4)
	var a, b, c, d int
	a = x + 1
	b = x - 1
	c = y + 1
	d = y - 1
	if a <= m-1 {
		res = append(res, arr[y][a])
	}
	if b >= 0 {
		res = append(res, arr[y][b])
	}
	if c <= n-1 {
		res = append(res, arr[c][x])
	}
	if d >= 0 {
		res = append(res, arr[d][x])
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])
		if i != len(res)-1 {
			fmt.Print(" ")
		}
	}
}
