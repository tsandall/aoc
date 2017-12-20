package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

const n = 256

var input1 = parseInputPart1("147,37,249,1,31,2,226,0,161,71,254,243,183,255,30,70")

var input2 = parseInputPart2("147,37,249,1,31,2,226,0,161,71,254,243,183,255,30,70")

func part1() int {
	list := make([]int, n)
	for i := range list {
		list[i] = i
	}
	hash(input1, list, 0, 0)
	return list[0] * list[1]
}

func part2() string {
	list := make([]int, n)
	for i := range list {
		list[i] = i
	}
	var skip, pos int
	for i := 0; i < 64; i++ {
		skip, pos = hash(input2, list, skip, pos)
	}
	var dense [16]int
	for densePos := range dense {
		for sparseOffset := 0; sparseOffset < 16; sparseOffset++ {
			sparsePos := (16 * densePos) + sparseOffset
			dense[densePos] ^= list[sparsePos]
		}
	}
	var hex string
	for i := range dense {
		hex += fmt.Sprintf("%02x", dense[i])
	}
	return hex
}

func hash(lengths []int, list []int, skip, pos int) (int, int) {
	for i := range lengths {
		hashOne(list, lengths[i], pos)
		pos += lengths[i] + skip
		skip++
	}
	return skip, pos
}

func hashOne(list []int, length, pos int) {
	for i := 0; i < (length / 2); i++ {
		a := mod(pos+i, n)
		b := mod(pos+length-i-1, n)
		tmp := list[a]
		list[a] = list[b]
		list[b] = tmp
	}
}

func mod(x, m int) int {
	y := x % m
	if y < 0 {
		return y + m
	}
	return y
}

func parseInputPart1(s string) (result []int) {
	for _, part := range strings.Split(s, ",") {
		i, err := strconv.ParseInt(part, 10, 64)
		if err != nil {
			panic(err)
		}
		result = append(result, int(i))
	}
	return
}

func parseInputPart2(s string) (result []int) {
	for i := range s {
		result = append(result, int(s[i]))
	}
	result = append(result, 17)
	result = append(result, 31)
	result = append(result, 73)
	result = append(result, 47)
	result = append(result, 23)
	return result
}
