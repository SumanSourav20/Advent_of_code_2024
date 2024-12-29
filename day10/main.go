package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type location struct {
	i, j int
}

func allHikingTrails(grid [][]int, i int, j int, peaks map[location]bool) int {
	if grid[i][j] == 9 {
		// if peaks[location{i, j}] {
		// 	return 0
		// }
		peaks[location{i, j}] = true
		return 1
	}

	count := 0
	num := grid[i][j]

	if i > 0 && grid[i-1][j] == num+1 {
		count += allHikingTrails(grid, i-1, j, peaks)
	}
	if i < len(grid)-1 && grid[i+1][j] == num+1 {
		count += allHikingTrails(grid, i+1, j, peaks)
	}
	if j > 0 && grid[i][j-1] == num+1 {
		count += allHikingTrails(grid, i, j-1, peaks)
	}
	if j < len(grid)-1 && grid[i][j+1] == num+1 {
		count += allHikingTrails(grid, i, j+1, peaks)
	}

	return count
}

func solve() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	scanner := bufio.NewScanner(file)
	var grid [][]int

	for scanner.Scan() {
		list := []int{}
		for _, ch := range scanner.Text() {
			num, _ := strconv.Atoi(string(ch))
			list = append(list, num)
		}
		grid = append(grid, list)
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				count += allHikingTrails(grid, i, j, map[location]bool{})
			}
		}
	}

	fmt.Println(count)
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
