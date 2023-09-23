package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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
