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
	input := parseInput(input)
	scanners := initScanners()
	var severity int

	for i := 0; i < len(input); i++ {
		// check if hit; increment severity if hit
		currLayer := input[i]
		if currLayer != nil && currLayer.scannerPos == 0 {
			severity += (i * currLayer.maxPos)
		}
		moveScanners(scanners, input)
	}

	return severity
}

func part2() int {
	input := parseInput(input)
	scanners := initScanners()
	for delay := 0; ; delay++ {
		var hit bool
		inputCopy := copyInput(input)
		scannersCopy := copyScanners(scanners)
		for i := 0; i < len(inputCopy); i++ {
			currLayer := inputCopy[i]
			if currLayer != nil && currLayer.scannerPos == 0 {
				hit = true
				break
			}
			moveScanners(scannersCopy, inputCopy)
		}
		if !hit {
			return delay
		}
		moveScanners(scanners, input)
	}
}

func initScanners() []bool {
	scanners := make([]bool, len(input))
	for i := range scanners {
		scanners[i] = true
	}
	return scanners
}

func copyInput(input []*layer) []*layer {
	cpy := make([]*layer, len(input))
	for i := range cpy {
		if input[i] != nil {
			tmp := *(input[i])
			cpy[i] = &tmp
		}
	}
	return cpy
}

func copyScanners(scanners []bool) []bool {
	cpy := make([]bool, len(scanners))
	copy(cpy, scanners)
	return cpy
}

func moveScanners(scanners []bool, input []*layer) {
	// move scanners to next spot
	for i, layer := range input {
		if layer != nil {
			if scanners[i] {
				layer.scannerPos++
				if layer.scannerPos == layer.maxPos {
					scanners[i] = false
					layer.scannerPos = layer.maxPos - 2
				}
			} else {
				layer.scannerPos--
				if layer.scannerPos < 0 {
					scanners[i] = true
					layer.scannerPos = 1
				}
			}
		}
	}
}

type layer struct {
	scannerPos int
	maxPos     int
}

func parseInput(s string) []*layer {
	var max int64
	unordered := map[int64]int64{}
	for _, line := range strings.Split(s, "\n") {
		flds := strings.Fields(line)
		layer, err := strconv.ParseInt(strings.TrimRight(flds[0], ":"), 10, 64)
		if err != nil {
			panic(err)
		}
		depth, err := strconv.ParseInt(flds[1], 10, 64)
		if err != nil {
			panic(err)
		}
		if layer > max {
			max = layer
		}
		unordered[layer] = depth
	}
	result := make([]*layer, max+1)
	for k, v := range unordered {
		result[int(k)] = &layer{
			scannerPos: 0,
			maxPos:     int(v),
		}
	}
	return result
}

var test = `0: 3
1: 2
4: 4
6: 4`

var input = `0: 4
1: 2
2: 3
4: 4
6: 8
8: 5
10: 6
12: 6
14: 10
16: 8
18: 6
20: 9
22: 8
24: 6
26: 8
28: 8
30: 12
32: 12
34: 12
36: 12
38: 10
40: 12
42: 12
44: 14
46: 8
48: 14
50: 12
52: 14
54: 14
58: 14
60: 12
62: 14
64: 14
66: 12
68: 12
72: 14
74: 18
76: 17
86: 14
88: 20
92: 14
94: 14
96: 18
98: 18`
