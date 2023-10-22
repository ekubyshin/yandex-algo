package L

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scanln(&n)
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	vals := make([]int64, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		if v, ok := strconv.ParseInt(s, 10, 64); ok == nil {
			vals = append(vals, v)
		}
	}
	scanner.Scan()
	k := scanner.Text()
	var ki int64
	var ok error
	if ki, ok = strconv.ParseInt(k, 10, 64); ok != nil {
		fmt.Println(-1)
		return
	}
	a, b := solv(ki, vals)
	fmt.Println(a, b)
}

func solv(p int64, elems []int64) (int, int) {
	from := 0
	a := solver(p, elems, from, len(elems))
	if a != -1 {
		from = a
	}
	b := solver(p*2, elems, from, len(elems))
	if a != -1 {
		a += 1
	}
	if b != -1 {
		b += 1
	}
	return a, b
}

func solver(x int64, elems []int64, left, right int) int {
	if right <= left {
		if left < len(elems) && elems[left] >= x {
			return left
		} else if right < len(elems) && elems[right] >= x {
			return right
		}
		return -1
	}
	mid := div(left+right, 2)
	if elems[mid] >= x {
		return solver(x, elems, left, mid)
	}
	return solver(x, elems, mid+1, right)
}

func div(a, b int) int {
	return int(math.Floor(float64(a) / float64(b)))
}
