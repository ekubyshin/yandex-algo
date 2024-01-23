package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

/*
К Васе в гости пришли одноклассники. Его мама решила угостить ребят печеньем.

Но не всё так просто. Печенья могут быть разного размера. А у каждого ребёнка есть фактор жадности —– минимальный размер печенья, которое он возьмёт. Нужно выяснить, сколько ребят останутся довольными в лучшем случае, когда они действуют оптимально.

Каждый ребёнок может взять не больше одного печенья.

Формат ввода
В первой строке записано n —– количество детей.

Во второй —– n чисел, разделённых пробелом, каждое из которых –— фактор жадности ребёнка. Это натуральные числа, не превосходящие 1000.

В следующей строке записано число m –— количество печенек.

Далее —– m натуральных чисел, разделённых пробелом —– размеры печенек. Размеры печенек не превосходят 1000.

Оба числа n и m не превосходят 10000.

Формат вывода
Нужно вывести одно число –— количество детей, которые останутся довольными

Пример 1
Ввод	Вывод
2
1 2
3
2 1 3
2
Пример 2
Ввод	Вывод
3
2 1 3
2
1 1
1
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	childrenCount, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	children := readArr(scanner.Text(), childrenCount)
	scanner.Scan()
	cookiesCount, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	cookies := readArr(scanner.Text(), cookiesCount)
	fmt.Println(solve(children, cookies))
}

func readArr(str string, count int) []int {
	res := make([]int, 0, count)
	for _, s := range strings.Split(str, " ") {
		n, _ := strconv.Atoi(s)
		res = append(res, n)
	}
	return res
}

func solve(children []int, cookies []int) (res int) {
	slices.Sort(cookies)
	slices.Sort(children)
	cookieNum := 0
	childNum := 0
	for cookieNum < len(cookies) && childNum < len(children) {
		if cookies[cookieNum] >= children[childNum] {
			res += 1
			childNum += 1
		}
		cookieNum += 1
	}
	return
}
