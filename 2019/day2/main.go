package main

import (
	"fmt"
	"strconv"
	"strings"
)

func crunch(x []int) int {

	for i := 0; i+3 < len(x); i += 4 {
		opcode := x[i]
		a := x[x[i+1]]
		b := x[x[i+2]]
		c := x[i+3]
		switch opcode {
		case 1:
			x[c] = a + b
		case 2:
			x[c] = a * b
		case 99:
			return x[0]
		default:
			panic("illegal opcode")
		}
	}

	panic("unreachable")
}

func part1(x []int) {
	fmt.Println(crunch(initializeProgram(x, 12, 2)))
}

func part2(input []int) {
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			result := crunch(initializeProgram(input, i, j))
			if result == 19690720 {
				fmt.Println(100*i + j)
			}
		}
	}
}

func initializeProgram(x []int, a, b int) []int {
	cpy := make([]int, len(x))
	copy(cpy, x)
	cpy[1] = a
	cpy[2] = b
	return cpy
}

func main() {
	part1(initializeProgram(input, 12, 2))
	part2(input)
}

func parseInput(s string) []int {
	var result []int
	for _, n := range strings.Split(s, ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}
	return result
}

var input = parseInput(`1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,5,19,23,2,9,23,27,1,6,27,31,1,31,9,35,2,35,10,39,1,5,39,43,2,43,9,47,1,5,47,51,1,51,5,55,1,55,9,59,2,59,13,63,1,63,9,67,1,9,67,71,2,71,10,75,1,75,6,79,2,10,79,83,1,5,83,87,2,87,10,91,1,91,5,95,1,6,95,99,2,99,13,103,1,103,6,107,1,107,5,111,2,6,111,115,1,115,13,119,1,119,2,123,1,5,123,0,99,2,0,14,0`)
