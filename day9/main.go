package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

func resultValues(text []byte) ([]int, int) {
	length := 0
	textSlice := []int{}

	for _, x := range text {
		num, _ := strconv.Atoi(string(x))
		textSlice = append(textSlice, num)
		length += num
	}

	return textSlice, length
}

func solve() {
	text, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	// var i, j int
	// if len(text)%2 == 0 {
	// 	j = len(text) - 2
	// } else {
	// 	j = len(text) - 1
	// }

	numSlice, _ := resultValues(text)
	result := []int{}

	// for _, x := range numSlice {
	// 	fmt.Print(x)
	// }
	// fmt.Println()

	for i, n := range numSlice {
		if i%2 == 0 {
			for range n {
				result = append(result, i/2)
			}
		} else {
			for range n {
				result = append(result, 0)
			}
		}
	}
	// fmt.Println(len(result), length)
	// for _, x := range result {
	// 	fmt.Print(x)
	// }

	newGaps := slices.Clone(numSlice)

	rIdx := len(result)
	// idx := len(newGaps[0])
	for i := len(numSlice) - 1; i >= 0; i-- {
		num := numSlice[i]
		rIdx = rIdx - num
		idx := 0
		if i%2 == 0 {
			// fmt.Println(i / 2)
			for j := 1; j < i; j += 2 {
				idx += newGaps[j-1]
				if newGaps[j] >= num {
					for itr := range num {
						result[idx+itr] = i / 2
						result[rIdx+itr] = 0
					}
					newGaps[j] = newGaps[j] - num
					newGaps[j-1] = newGaps[j-1] + num
					break
				}
				idx += newGaps[j]
			}

			// for _, x := range newGaps {
			// 	fmt.Print(x)
			// }
			// fmt.Println()
		}
	}

	total := 0
	for i, x := range result {
		// fmt.Print(x)
		total += i * x
	}
	fmt.Println(total)

	// itr, _ := strconv.Atoi(string(text[j]))
	// total := 0
	// // index := 0
	// for i < j-1 {
	// 	x, _ := strconv.Atoi(string(text[i]))

	// 	for range x {
	// 		result = append(result, i/2)
	// 		// total += index * i / 2
	// 		// index++
	// 	}

	// 	x, _ = strconv.Atoi(string(text[i+1]))

	// 	for range x {
	// 		if itr == 0 {
	// 			j -= 2
	// 			if j <= i {
	// 				break
	// 			}
	// 			itr, _ = strconv.Atoi(string(text[j]))
	// 		}

	// 		result = append(result, j/2)
	// 		// total += index * j / 2
	// 		// index++
	// 		itr--
	// 	}
	// 	i += 2
	// }

	// for itr > 0 {
	// 	result = append(result, j/2)
	// 	// total += index * j / 2
	// 	// index++
	// 	itr--
	// }

	// // total := 0
	// // for i, r := range result {
	// // 	total += i * r
	// // }

	// fmt.Println(total)
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
