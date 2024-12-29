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

func ordered(hashmap map[int][]int, numbers []int) bool {
	for i, num := range numbers {
		v, _ := hashmap[num]
		for j := 0; j < i; j++ {
			contain := slices.Contains(v, numbers[j])
			if contain {
				return false
			}
		}
	}
	return true
}

func order(hashmap map[int][]int, numbers []int) {
	for i, num := range numbers {
		v, _ := hashmap[num]
		for j := i - 1; j >= 0; j-- {
			contain := slices.Contains(v, numbers[j])
			if contain {
				temp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = temp
			} else {
				break
			}
		}
	}
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

	hashmap := make(map[int][]int)
	total := 0
	otherTotal := 0

	for scanner.Scan() {
		line := scanner.Text()

		if strings.ContainsRune(line, '|') {
			numbers := strings.Split(line, "|")

			num1, err := strconv.Atoi(numbers[0])
			num2, err := strconv.Atoi(numbers[1])

			if err != nil {
				log.Fatal("Error opening the file : %w", err)
			}

			v, ok := hashmap[num1]
			if ok == false {
				hashmap[num1] = []int{num2}
			} else {
				v = append(v, num2)
				hashmap[num1] = v
			}
		} else {
			numbers := strings.Split(line, ",")
			var numbersInt []int
			for _, num := range numbers {
				intNum, _ := strconv.Atoi(num)
				numbersInt = append(numbersInt, intNum)
			}

			if ordered(hashmap, numbersInt) {
				total += numbersInt[len(numbersInt)/2]
				// fmt.Println(numbersInt[len(numbersInt)/2])
			} else {
				order(hashmap, numbersInt)
				otherTotal += numbersInt[len(numbersInt)/2]
			}

		}
	}
	// for k, v := range hashmap {
	// 	fmt.Printf("%d:", k)
	// 	for _, i := range v {
	// 		fmt.Printf("%d ", i)
	// 	}
	// }

	fmt.Println(total, otherTotal)

}
