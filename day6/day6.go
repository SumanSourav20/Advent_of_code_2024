package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type visit struct {
	ic, jc int
	i, j   int
}

func isLoopTraversal(i int, j int, grid [][]byte) bool {
	ic := -1
	jc := 0
	// count := 0
	var loop bool
	visited := map[visit]bool{}

	for {
		// grid[i][j] = 'X'
		if (ic == -1 && i == 0) || (jc == 1 && j == len(grid[0])-1) || (ic == 1 && i == len(grid)-1) || (jc == -1 && j == 0) {
			// return count
			loop = false
			break
		}

		if visited[visit{ic, jc, i + ic, j + jc}] {
			loop = true
			break
		}

		if grid[i+ic][j+jc] == '#' {
			visited[visit{ic, jc, i + ic, j + jc}] = true
			switch {
			case ic == -1:
				ic = 0
				jc = 1
			case ic == 1:
				ic = 0
				jc = -1
			case jc == -1:
				ic = -1
				jc = 0
			case jc == 1:
				ic = 1
				jc = 0
			}
		}
		if grid[i+ic][j+jc] != '#' {
			i = i + ic
			j = j + jc
		}
		// count++
	}

	for x := range grid {
		for y := range grid[x] {
			if grid[x][y] == '1' || grid[x][y] == '2' || grid[x][y] == '3' || grid[x][y] == '4' {
				grid[x][y] = '#'
			}
			// fmt.Printf("%c", grid[x][y])
		}
		// fmt.Println()
	}

	return loop
}

func countLoops(grid [][]byte, i int, j int) {
	count := 0
	for x := range grid {
		for y := range grid[x] {
			if !(i == x && j == y) && grid[x][y] != '#' {
				grid[x][y] = '#'
				if isLoopTraversal(i, j, grid) {
					count++
					// fmt.Print("O")
				} else {
					// fmt.Print(".")
				}
				grid[x][y] = '.'
			} else if i == x && j == y {
				// fmt.Print("^")
			} else {
				// fmt.Print("#")
			}
		}
		// fmt.Println()
	}
	fmt.Println(count)
}
func print(grid [][]byte) {
	for x := range grid {
		for y := range grid[x] {
			fmt.Printf("%c", grid[x][y])
		}
		fmt.Println()
	}
}
func countDistincPath(grid [][]byte) {
	for i, g := range grid {
		for j, ch := range g {
			if ch == '^' {
				grid[i][j] = '.'
				fmt.Println("found ^")
				countLoops(grid, i, j)
				// print(grid)
				// print(grid)
				// fmt.Print(isLoopTraversal(i, j, grid))
			}
		}
	}

	// count := 0
	// for _, g := range grid {
	// 	for _, ch := range g {
	// 		// fmt.Printf("%c", grid[i][j])
	// 		if ch == 'X' {
	// 			count++
	// 		}
	// 	}
	// 	// fmt.Println()
	// }
	// fmt.Println(count)

}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("No input file specified")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Error opening the file : %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var grid [][]byte

	for scanner.Scan() {
		line := scanner.Text()
		var arr []byte
		for _, ch := range line {
			arr = append(arr, byte(ch))
		}

		grid = append(grid, arr)
	}

	countDistincPath(grid)

}
