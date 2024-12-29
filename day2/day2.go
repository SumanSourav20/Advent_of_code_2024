package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func strictlyIncreasing(list []string) bool {
	badLevelAlowed := true
	for i := 1; i < len(list); i++ {
		num, err := strconv.Atoi(list[i])
		prev, err := strconv.Atoi(list[i-1])
		// fmt.Println(num, prev)

		if err != nil {
			log.Fatal("unable to convert string word to int")
		}

		if num <= prev || num > prev+3 {
			if badLevelAlowed {
				list[i] = list[i-1]
				badLevelAlowed = false
				continue
			}
			return false
		}

	}
	return true
}

func strictlyDecreasing(list []string) bool {
	badLevelAlowed := true
	for i := 1; i < len(list); i++ {
		num, err := strconv.Atoi(list[i])
		prev, err := strconv.Atoi(list[i-1])
		// fmt.Println(num, prev)

		if err != nil {
			log.Fatal("unable to convert string word to int")
		}

		if num >= prev || num < prev-3 {
			if badLevelAlowed {
				list[i] = list[i-1]
				badLevelAlowed = false
				continue
			}
			return false
		}
	}
	return true
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

	safeRecords := 0

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Fields(line)
		reverse := slices.Clone(list)
		slices.Reverse(reverse)

		if strictlyIncreasing(slices.Clone(list)) || strictlyDecreasing(slices.Clone(list)) || strictlyDecreasing(slices.Clone(reverse)) || strictlyIncreasing(slices.Clone(reverse)) {
			fmt.Println("Yes")
			safeRecords++
		} else {
			fmt.Println("No")
		}
	}
	fmt.Println(safeRecords)
}
