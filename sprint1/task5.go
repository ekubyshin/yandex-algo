package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)
	max := 0
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	accum := 0
	mword := ""
	for accum < n {
		scanner.Scan()
		w := scanner.Text()
		l := len(w)
		if l > max {
			max = l
			mword = w
		}
		accum += l + 1
	}
	fmt.Println(mword)
	fmt.Println(max)
}
