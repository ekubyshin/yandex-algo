package main

//91183578

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var k int
	fmt.Scanln(&k)
	m := readArray()
	out := 0
	max := k * 2
	for _, v := range m {
		if v <= max {
			out += 1
		}
	}
	fmt.Println(out)
}

func readArray() map[int]int {
	reader := bufio.NewReader(os.Stdin)
	max := 4 * 4
	m := make(map[int]int)
	i := 0
	for {
		if b, ok := reader.ReadByte(); ok == nil {
			if b == 10 {
				continue
			}
			num := 0
			if b >= '1' && b <= '9' {
				num, _ = strconv.Atoi(string(b))
				if v, ok := m[num]; ok {
					m[num] = v + 1
				} else {
					m[num] = 1
				}
			}
			i += 1
			if i == max {
				break
			}
		} else {
			break
		}
	}
	return m
}
