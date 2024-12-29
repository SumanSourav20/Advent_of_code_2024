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

func recursiveMatch(test int, nums []string, i int, result int) bool {
	if i >= len(nums) {
		return result == test
	}

	num, _ := strconv.Atoi(nums[i])
	res := strconv.Itoa(result)
	res = res + nums[i]
	resNum, _ := strconv.Atoi(res)
	return recursiveMatch(test, nums, i+1, result+num) || recursiveMatch(test, nums, i+1, result*num) || recursiveMatch(test, nums, i+1, resNum)
}

func match() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")

		test, err := strconv.Atoi(lineSplit[0])

		if err != nil {
			log.Fatal("Error converting to int:", err)
		}

		numStrings := strings.Fields(lineSplit[1])

		// nums := []int{}
		// for _, num := range numStrings {
		// 	n, _ := strconv.Atoi(num)
		// 	nums = append(nums, n)
		// }
		result, _ := strconv.Atoi(numStrings[0])
		if recursiveMatch(test, numStrings, 1, result) {
			total += test
		}
	}

	fmt.Println(total)
}

func main() {
	start := time.Now()
	match()
	duration := time.Since(start)
	fmt.Println(duration)
}
