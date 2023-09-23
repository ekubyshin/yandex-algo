package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
