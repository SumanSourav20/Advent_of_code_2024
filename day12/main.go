package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func perimeterContribution(grid [][]int, i, j int, ch int) int {
	count := 0
	if i == 0 {
		if j == len(grid[i])-1 || grid[i][j+1] != ch {
			count++
		}
	} else if grid[i-1][j] != ch {
		if j == len(grid[i])-1 || grid[i][j+1] != ch || grid[i-1][j+1] == ch {
			count++
		}
	}

	if i == len(grid)-1 {
		if j == len(grid[i])-1 || grid[i][j+1] != ch {
			count++
		}
	} else if grid[i+1][j] != ch {
		if j == len(grid[i])-1 || grid[i][j+1] != ch || grid[i+1][j+1] == ch {
			count++
		}
	}

	if j == 0 {
		if i == len(grid)-1 || grid[i+1][j] != ch {
			count++
		}
	} else if grid[i][j-1] != ch {
		if i == len(grid)-1 || grid[i+1][j] != ch || grid[i+1][j-1] == ch {
			count++
		}
	}

	if j == len(grid[i])-1 {
		if i == len(grid)-1 || grid[i+1][j] != ch {
			count++
		}
	} else if grid[i][j+1] != ch {
		if i == len(grid)-1 || grid[i+1][j] != ch || grid[i+1][j+1] == ch {
			count++
		}
	}

	return count
}

type price struct {
	count, perimeter int
}

type direction struct {
	i, j int
}

func dfs(grid [][]byte, labels [][]int, ch byte, i, j int, label int) {
	directions := []direction{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[i]) {
		return
	}

	if grid[i][j] != ch || labels[i][j] != 0 {
		return
	}
	labels[i][j] = label
	for _, d := range directions {
		dfs(grid, labels, ch, i+d.i, j+d.j, label)
	}
}

func visit(grid [][]byte, labels [][]int) {
	id := 1
	for i := range grid {
		for j := range grid[i] {
			if labels[i][j] == 0 {
				dfs(grid, labels, grid[i][j], i, j, id)
				id++
			}
		}
	}

}
func calculateTotalPrice(grid [][]byte) int {
	prices := map[int]price{}

	rows := len(grid)
	cols := len(grid[0])

	labels := make([][]int, rows)
	for i := range labels {
		l := make([]int, cols)
		labels[i] = l
	}

	visit(grid, labels)

	total := 0
	for i, row := range labels {
		for j, ch := range row {
			p, _ := prices[ch]
			p.perimeter += perimeterContribution(labels, i, j, ch)
			p.count++

			prices[ch] = p
		}
	}

	for ch := range prices {
		// fmt.Println(prices[ch].count, prices[ch].perimeter, ch)
		total += prices[ch].count * prices[ch].perimeter
	}

	return total
}

func solve() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	grid := [][]byte{}

	for scanner.Scan() {
		chars := scanner.Text()

		list := []byte{}
		for _, ch := range chars {
			list = append(list, string(ch)[0])
		}

		grid = append(grid, list)
	}

	totalPrice := calculateTotalPrice(grid)
	fmt.Println(totalPrice)
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
