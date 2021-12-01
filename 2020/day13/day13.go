package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	var buses [][2]int

	for i, s := range strings.Split(strings.Split(input, "\n")[1], ",") {
		if s == "x" {
			continue
		}
		id, _ := strconv.Atoi(s)
		buses = append(buses, [2]int{i, id})
	}

	t := 0
	step := 1

	for _, bus := range buses {
		offset, period := bus[0], bus[1]
		for (t+offset)%period != 0 {
			t += step
		}
		fmt.Println("found:", "bus:", period, "t:", t, "step:", step)
		step *= period
	}

	fmt.Println(t)
}

const input = `1000507
29,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,631,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,x,13,19,x,x,x,23,x,x,x,x,x,x,x,383,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17`
