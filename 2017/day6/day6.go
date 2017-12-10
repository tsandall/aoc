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

func part1() int {
	seen := map[banks]struct{}{}
	var count int
	curr := input
	for {
		if _, ok := seen[curr]; ok {
			break
		}
		seen[curr] = struct{}{}
		pos := curr.Max()
		curr = curr.Redistribute(pos)
		count++
	}
	return count
}

func part2() int {
	seen := map[banks]int{}
	var cycle int
	curr := input
	for {
		if exist, ok := seen[curr]; ok {
			return cycle - exist
		}
		seen[curr] = cycle
		pos := curr.Max()
		curr = curr.Redistribute(pos)
		cycle++
	}
}

type banks [16]int

// Max returns the position of the bank with the highest population. Ties are
// broken by picking the bank with the lowest position.
func (b banks) Max() int {
	maxValue := b[0]
	maxPos := 0
	for i := 1; i < len(b); i++ {
		if b[i] > maxValue {
			maxPos = i
			maxValue = b[i]
		}
	}
	return maxPos
}

// Redistribute takes the population and redistributes it across the banks.
func (b banks) Redistribute(pos int) banks {
	cpy := b
	pop := cpy[pos]
	cpy[pos] = 0
	i := pos
	for pop > 0 {
		i = (i + 1) % len(cpy)
		cpy[i]++
		pop--
	}
	return cpy
}

var input = parseInput(`11 11 13 7 0 15 5 5 4 4 1 1 7 1 15 11`)

func parseInput(s string) (result banks) {
	for i, pop := range strings.Fields(s) {
		i64, err := strconv.ParseInt(pop, 10, 64)
		if err != nil {
			panic(err)
		}
		result[i] = int(i64)
	}
	return
}
