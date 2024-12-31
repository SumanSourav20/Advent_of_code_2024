package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func print(grid [][]byte) {
	for _, g := range grid {
		for _, ch := range g {
			fmt.Printf("%c", ch)
		}
		fmt.Println()
	}
}

func readInput(filename string) ([][]byte, string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	scanner := bufio.NewScanner(file)

	grid := [][]byte{}

	for scanner.Scan() {
		chars := scanner.Text()

		if len(chars) == 0 {
			break
		}

		list := []byte{}
		for _, ch := range chars {
			list = append(list, string(ch)[0])
		}

		grid = append(grid, list)
	}

	directions := ""
	for scanner.Scan() {
		directions += scanner.Text()
	}

	return grid, directions
}

func move(grid [][]byte, i, j int, ic, jc int) bool {
	if grid[i+ic][j+jc] == '#' {
		return false
	}
	if grid[i+ic][j+jc] == '.' {
		grid[i+ic][j+jc] = grid[i][j]
		grid[i][j] = '.'
		return true
	}
	if move(grid, i+ic, j+jc, ic, jc) {
		grid[i+ic][j+jc] = grid[i][j]
		grid[i][j] = '.'
		return true
	}
	return false
}

func blocked(grid [][]byte, i int, indexes []int, ic int) bool {
	// for _, j := range indexes {
	// 	fmt.Println(i, j, ic)
	// }
	for _, in := range indexes {
		if grid[i+ic][in] == '#' {
			return true
		}
	}
	return false
}

func nextBoxes(grid [][]byte, i int, indexes []int, ic int) []int {
	result := []int{}
	for _, in := range indexes {
		if grid[i+ic][in] == '[' {
			result = append(result, in)
			result = append(result, in+1)
		} else if grid[i+ic][in] == ']' {
			if len(result) == 0 || !(result[len(result)-1] == in) {
				result = append(result, in-1)
				result = append(result, in)
			}
		}
	}

	return result
}

func moveUpDown(grid [][]byte, i int, indexes []int, ic int) bool {
	if blocked(grid, i, indexes, ic) {
		return false
	}
	next := nextBoxes(grid, i, indexes, ic)
	// fmt.Println(len(next))
	if len(next) == 0 {
		for _, in := range indexes {
			grid[i+ic][in] = grid[i][in]
			grid[i][in] = '.'
		}
		return true
	}

	if moveUpDown(grid, i+ic, next, ic) {
		for _, in := range indexes {
			grid[i+ic][in] = grid[i][in]
			grid[i][in] = '.'
		}
		return true
	}

	return false
}

func moveAll(grid [][]byte, directions string) [][]byte {
	for i, g := range grid {
		for j, ch := range g {
			if ch == '@' {
				for _, ch := range directions {
					// fmt.Println(i, j)
					// print(grid)
					// fmt.Printf("%c", ch)
					switch string(ch) {
					case "<":
						if move(grid, i, j, 0, -1) {
							j -= 1
						}
					case ">":
						if move(grid, i, j, 0, 1) {
							j += 1
						}
					case "^":
						if moveUpDown(grid, i, []int{j}, -1) {
							i -= 1
						}
					case "v":
						if moveUpDown(grid, i, []int{j}, 1) {
							i += 1
						}
					}
				}

				return grid

			}
		}
	}
	return grid
}

func result(grid [][]byte) int {
	total := 0
	for i, g := range grid {
		for j, ch := range g {
			if ch == '[' {
				total += 100*i + j
			}
		}
	}
	return total
}

func part2Grid(grid [][]byte) [][]byte {
	output := [][]byte{}
	for _, g := range grid {
		result := []byte{}
		for _, ch := range g {
			if ch == '@' {
				result = append(result, '@')
				result = append(result, '.')
			} else if ch == 'O' {
				result = append(result, '[')
				result = append(result, ']')
			} else {
				result = append(result, ch)
				result = append(result, ch)
			}
		}
		output = append(output, result)
	}
	return output
}
func solve() {
	grid, directions := readInput("input.txt")
	grid = part2Grid(grid)
	print(grid)
	fmt.Println(len(directions))
	grid = moveAll(grid, directions)
	fmt.Println()
	print(grid)
	fmt.Println(result(grid))
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
