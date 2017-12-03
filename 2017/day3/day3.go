package main

import "fmt"

const input = 347991

func spiral(f func(int, int, int) bool) {
	var x, y int
	var phase bool
	radius := 0

	for n := 1; ; n++ {
		if phase {
			if x < radius {
				x++
			} else if y < radius {
				y++
			}
			if x == radius && y == radius {
				phase = !phase
			}
		} else {
			if x > -radius {
				x--
			} else if y > -radius {
				y--
			}
			if x == -radius && y == -radius {
				phase = !phase
			}
		}
		if f(n, x, y) {
			return
		}
		if x == radius && y == -radius {
			radius++
		}
	}
}

func count1() int {
	var rx, ry int
	spiral(func(n, x, y int) bool {
		rx, ry = x, y
		return n == input
	})

	if rx < 0 {
		rx = -rx
	}
	if ry < 0 {
		ry = -rx
	}
	return rx + ry
}

func count2() int {

	sumGrid := map[[2]int]int{
		{0, 0}: 1,
	}

	var result int

	spiral(func(n, x, y int) bool {
		sum := 0
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				k := [2]int{x + i, y + j}
				if val, ok := sumGrid[k]; ok {
					sum += val
				}
			}
		}
		sumGrid[[2]int{x, y}] = sum
		if sum > input {
			result = sum
			return true
		}
		return false
	})

	return result
}

func main() {
	fmt.Println(count1())
	fmt.Println(count2())
}
