package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	var list []string

	for scanner.Scan() {
		line := scanner.Text()
		list = append(list, line)
	}

	// count := 0
	// change := []int{0, 1, -1}

	// for i := range len(list) {
	// 	for j := range len(list[i]) {
	// 		for _, ic := range change {
	// 			for _, jc := range change {
	// 				if i+ic*3 < len(list) && i+ic*3 >= 0 && j+jc*3 < len(list[i]) && j+jc*3 >= 0 {
	// 					if list[i][j] == 'X' && list[i+ic][j+jc] == 'M' && list[i+ic*2][j+jc*2] == 'A' && list[i+ic*3][j+jc*3] == 'S' {
	// 						count++
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	count := 0
	change := []int{1, -1}
	for i := 1; i < len(list)-1; i++ {
		for j := 1; j < len(list)-1; j++ {
			for _, ic := range change {
				for _, jc := range change {
					if list[i][j] == 'A' && list[i+ic][j+jc] == 'M' && list[i-ic][j-jc] == 'S' && ((list[i-ic][j+jc] == 'M' && list[i+ic][j-jc] == 'S') || (list[i+ic][j-jc] == 'M' && list[i-ic][j+jc] == 'S')) {
						count++
					}
				}
			}
		}
	}

	fmt.Println(count / 2)
}
