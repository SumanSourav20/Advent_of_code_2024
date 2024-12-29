package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type location struct {
	i, j int
}

func (x location) add(y location) location {
	return location{x.i + y.i, x.j + y.j}
}

func (x location) sub(y location) location {
	return location{x.i - y.i, x.j - y.j}
}

func uniqueLocations(grid [][]byte) map[byte][]location {
	uniqueAntinaLocations := map[byte][]location{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '.' {
				continue
			}
			list, ok := uniqueAntinaLocations[grid[i][j]]
			if ok {
				list = append(list, location{i, j})
				uniqueAntinaLocations[grid[i][j]] = list
			} else {
				uniqueAntinaLocations[grid[i][j]] = []location{{i, j}}
			}
		}
	}
	return uniqueAntinaLocations
}

// func markLocation(i int, j int) int {
// 	return i * len[g]
// }

func uniqueAntinodes(grid [][]byte) int {
	markGrid := make([]bool, len(grid)*len(grid[0]))
	// fmt.Println(len(markGrid))
	uniqueAntinaLocations := uniqueLocations(grid)
	for antina := range uniqueAntinaLocations {
		n := len(uniqueAntinaLocations[antina])
		// fmt.Println(n)
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				diff := uniqueAntinaLocations[antina][i].sub(uniqueAntinaLocations[antina][j])

				antinode1 := uniqueAntinaLocations[antina][i]
				// fmt.Print(antinode1)
				for antinode1.i >= 0 && antinode1.i < len(grid) && antinode1.j >= 0 && antinode1.j < len(grid[0]) {
					loc := antinode1.i*len(grid[0]) + antinode1.j
					markGrid[loc] = true
					antinode1 = antinode1.add(diff)
				}
				antinode2 := uniqueAntinaLocations[antina][j]
				// fmt.Println(antinode2)
				for antinode2.i >= 0 && antinode2.i < len(grid) && antinode2.j >= 0 && antinode2.j < len(grid[0]) {
					loc := antinode2.i*len(grid[0]) + antinode2.j
					markGrid[loc] = true
					antinode2 = antinode2.sub(diff)
				}
			}
		}
	}

	count := 0

	for i := range grid {
		for j := range grid[i] {
			if markGrid[i*len(grid[0])+j] {
				fmt.Print("#")
				count++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

	// for i := 0; i < len(grid)*len(grid[0]); i++ {
	// 	if markGrid[i] {
	// 		count++
	// 	}
	// }

	return count
}

func solve() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	scanner := bufio.NewScanner(file)
	var grid [][]byte

	for scanner.Scan() {
		list := []byte{}
		for _, ch := range scanner.Text() {
			list = append(list, byte(ch))
		}
		grid = append(grid, list)
	}

	for i := range grid {
		for j := range grid[i] {
			fmt.Printf("%c", grid[i][j])
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println(uniqueAntinodes(grid))
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
