package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type location struct {
	i, j int
}

func isEvenDigits(num int) bool {
	count := 0
	for num != 0 {
		num = num / 10
		count++
	}
	return count%2 == 0
}

func split(num int) (int, int) {
	numStr := strconv.Itoa(num)
	halflen := len(numStr) / 2
	firstHalf := numStr[:halflen]
	secondHalf := numStr[halflen:]

	firstNum, _ := strconv.Atoi(firstHalf)
	secondNum, _ := strconv.Atoi(secondHalf)

	return firstNum, secondNum
}

type numDepth struct {
	num, i int
}

func totalBranching(num int, i int, cache map[numDepth]int) int {
	if i <= 0 {
		return 1
	}

	count, ok := cache[numDepth{num, i}]
	if ok {
		return count
	}

	if num == 0 {
		return totalBranching(1, i-1, cache)
	} else if isEvenDigits(num) {
		a, b := split(num)
		count = totalBranching(a, i-1, cache) + totalBranching(b, i-1, cache)
		cache[numDepth{num, i}] = count
		return count
	} else {
		count = totalBranching(num*2024, i-1, cache)
		cache[numDepth{num, i}] = count
		return count
	}
}

func solve() {
	text, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	input := string(text)

	nums := strings.Split(input, " ")
	numsInt := []int{}
	for _, n := range nums {
		x, _ := strconv.Atoi(n)
		numsInt = append(numsInt, x)
	}

	count := 0
	for _, n := range numsInt {
		count += totalBranching(n, 75, map[numDepth]int{})
	}

	fmt.Println(count)
}

func main() {
	start := time.Now()
	solve()
	duration := time.Since(start)
	fmt.Println(duration)
}
