package main

import (
	"fmt"
	"strconv"
	"strings"
)

func nth(n int) int {
	turns := map[int][2]int{}
	var last int

	for i, str := range strings.Split(input, ",") {
		t := i + 1
		var err error
		last, err = strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		turns[last] = [2]int{t, 0}
	}

	for t := len(turns) + 1; t <= n; t++ {
		pair := turns[last]
		if pair[1] == 0 {
			last = 0
		} else {
			last = pair[0] - pair[1]
		}
		turns[last] = [2]int{t, turns[last][0]}
	}

	return last
}

func main() {

	fmt.Println(nth(2020))
	fmt.Println(nth(30 * 1000 * 1000))
}

// const input = "0,3,6"

const input = `8,0,17,4,1,12`
