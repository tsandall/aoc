package main

import (
	"fmt"
	"strings"
)

type pos struct {
	x int
	y int
	z int
	w int
}

func countActiveNeighbours(active map[pos]struct{}, p pos) int {
	var n int
	for _, q := range neighbours(p) {
		if _, ok := active[q]; ok {
			n++
		}
	}
	return n
}

func neighbours(p pos) []pos {
	var result []pos
	deltas := [3]int{-1, 0, 1}
	for _, i := range deltas {
		for _, j := range deltas {
			for _, k := range deltas {
				for _, l := range deltas {
					if i == 0 && j == 0 && k == 0 && l == 0 {
						continue
					}
					result = append(result, pos{
						x: p.x + i,
						y: p.y + j,
						z: p.z + k,
						w: p.w + l,
					})
				}
			}
		}
	}
	return result
}

func main() {

	space := parseSpace(input)

	active := map[pos]struct{}{}

	for w := range space {
		for z := range space[w] {
			for y := range space[w][z] {
				for x := range space[w][z][y] {
					if space[w][z][y][x] == "#" {
						active[pos{x, y, z, w}] = struct{}{}
					}
				}
			}
		}
	}

	for i := 1; i <= 6; i++ {

		toInactive := map[pos]struct{}{}

		for p := range active {
			n := countActiveNeighbours(active, p)
			if n != 2 && n != 3 {
				toInactive[p] = struct{}{}
			}
		}

		toActive := map[pos]struct{}{}

		for p := range active {
			for _, q := range neighbours(p) {
				_, isActive := active[q]
				if !isActive && countActiveNeighbours(active, q) == 3 {
					toActive[q] = struct{}{}
				}
			}
		}

		for p := range toActive {
			active[p] = struct{}{}
		}

		for p := range toInactive {
			delete(active, p)
		}

		fmt.Printf("#%d active: %v\n", i, len(active))
	}

}

func parseSpace(s string) [][][][]string {
	var plane [][]string
	for _, l := range strings.Split(s, "\n") {
		line := strings.Split(l, "")
		plane = append(plane, line)
	}
	return [][][][]string{[][][]string{plane}}
}

const smallInput = `.#.
..#
###`

const input = `....#...
.#..###.
.#.#.###
.#....#.
...#.#.#
#.......
##....#.
.##..#.#`

/*

z=-2
.......
.......
..@@...
..@@@..
.......
.......
.......

z=-1
..@....
...@...
@......
.....@@
.@...@.
..@.@..
...@...

z=0
...@...
.......
@......
.......
.....@@
.@@.@..
...@...

z=1
..@....
...@...
@......
.....@@
.@...@.
..@.@..
...@...

z=2
.......
.......
..@@...
..@@@..
.......
.......
.......

*/
