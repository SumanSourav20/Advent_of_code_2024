package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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
	var list1, list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		list := strings.Fields(line)
		val1, err := strconv.Atoi(list[0])
		val2, err := strconv.Atoi(list[1])

		if err != nil {
			log.Fatal("unalbe to convert string word to int")
		}

		list1 = append(list1, val1)
		list2 = append(list2, val2)
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var total float64
	for i := range list1 {
		fmt.Printf("%d %d\n", list1[i], list2[i])
		total += math.Abs(float64(list1[i] - list2[i]))
	}
	fmt.Printf("%d", int(total))
}
