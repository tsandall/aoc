package main

import (
	"fmt"
	"strconv"
)

func meetsCriteria(bs [6]byte) (bool, bool) {

	var count [10]int

	for i := 0; i < 6; i++ {

		if i < 5 && bs[i] > bs[i+1] {
			return false, false
		}

		count[bs[i]]++
	}

	var case1, case2 bool

	for _, c := range count {
		if c > 1 {
			case1 = true
			if c == 2 {
				case2 = true
			}
		}
	}

	return case1, case2
}

func add1(bs [6]byte) [6]byte {

	pos := 5

	for pos >= 0 {
		v := bs[pos] + 1
		if v <= 9 {
			bs[pos] = v
			return bs
		}
		bs[pos] = 0
		pos--
	}

	return bs
}

func run(min, max int) {

	var bs [6]byte

	s := strconv.Itoa(min)
	if len(s) != 6 {
		panic(s)
	}

	for i := range s {
		bs[i] = s[i] - '0'
	}

	var part1, part2 int

	for i := min; i <= max; i++ {
		ok1, ok2 := meetsCriteria(bs)
		if ok1 {
			part1++
		}
		if ok2 {
			part2++
		}
		bs = add1(bs)
	}

	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2)
}

func main() {
	run(264793, 803935)
}
