package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type point struct {
	x, y int
}

func makePoint(p string) point {
	pl := strings.Split(p, ",")
	px, _ := strconv.Atoi(pl[0])
	py, _ := strconv.Atoi(pl[1])

	return point{px, py}
}

func splitInputLine(line string) (point, point) {
	list := strings.Split(line, " ")
	p := strings.Split(list[0], "=")[1]
	v := strings.Split(list[1], "=")[1]

	return makePoint(p), makePoint(v)
}

func newPosition(p, v point, n, m, seconds int) point {
	newX := 0
	newY := 0
	if v.x < 0 {
		newX = (n + (p.x+v.x*seconds)%n) % n
	} else {
		newX = (p.x + v.x*seconds) % n
	}
	if v.y < 0 {
		newY = (m + (p.y+v.y*seconds)%m) % m
	} else {
		newY = (p.y + v.y*seconds) % m

	}
	return point{newX, newY}
}

// func quadrent(n, m int, position, velocity point) (int, bool) {
// 	newP := newPosition(position, velocity, n, m, 100)

// 	if newP.x < n/2 {
// 		if newP.y < m/2 {
// 			return 1, true
// 		} else if newP.y > (m-1)/2 {
// 			return 3, true
// 		}
// 	} else if newP.x > (n-1)/2 {
// 		if newP.y < m/2 {
// 			return 2, true
// 		} else if newP.y > (m-1)/2 {
// 			return 4, true
// 		}
// 	}
// 	return 0, false
// }

type input struct {
	p, v point
}

func print(grid [][]bool) {
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func findConsecutive(grid [][]bool) bool {
outer:
	for i := 0; i < len(grid)-10; i++ {
		for j := i; j < i+10; j++ {
			if !grid[i][j] {
				continue outer
			}
		}
		return true
	}
	return false
}

func solve() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	// total := 1
	// list := []int{0, 0, 0, 0}
	list := []input{}
	for scanner.Scan() {
		line := scanner.Text()
		p, v := splitInputLine(line)
		list = append(list, input{p, v})
		// quad, ok := quadrent(101, 103, p, v)
		// if ok {
		// 	list[quad-1] = list[quad-1] + 1
		// }

	}

	for i := 1; i <= 10000; i++ {
		grid := [][]bool{}
		for range 103 {
			row := make([]bool, 101)
			grid = append(grid, row)
		}

		for _, l := range list {
			p := newPosition(l.p, l.v, 101, 103, i)
			grid[p.y][p.x] = true
		}

		if findConsecutive(grid) {
			fmt.Println(i)
			print(grid)
		}
	}

	// for _, v := range list {
	// 	total *= v
	// }
	// fmt.Println(total)
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
