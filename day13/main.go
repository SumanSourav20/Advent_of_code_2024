package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func (p point) add(p2 point) point {
	return point{p.x + p2.x, p.y + p2.y}
}

func (p point) sub(p2 point) point {
	return point{p.x - p2.x, p.y - p2.y}
}

func xy(line string) (int, int) {
	a := strings.Split(line, ":")[1]
	var ax, ay string
	if strings.Index(a, "+") != -1 {
		ax = strings.Split(strings.Split(a, ",")[0], "+")[1]
		ay = strings.Split(strings.Split(a, ",")[1], "+")[1]
	} else {
		ax = strings.Split(strings.Split(a, ",")[0], "=")[1]
		ay = strings.Split(strings.Split(a, ",")[1], "=")[1]
	}

	x, _ := strconv.Atoi(ax)
	y, _ := strconv.Atoi(ay)

	return x, y
}

func HCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minCostOfReaching(destination point, stepA point, stepB point, costA int, costB int) int {
	// location := point{0, 0}
	// bCount := 0
	// fmt.Println(stepB.x, stepB.y)
	// for location.x < destination.x && location.y < destination.y {
	// 	location = location.add(stepB)
	// 	bCount++
	// }
	// bCount = min(destination.x/stepB.x, destination.y/stepB.y) + 1
	// location = point{stepB.x * bCount, stepB.y * bCount}

	// for bCount >= 0 {
	// 	location = location.sub(stepB)
	// 	bCount--
	// 	diff := destination.sub(location)
	// 	if diff.x%stepA.x == 0 && diff.y%stepA.y == 0 && diff.x/stepA.x == diff.y/stepA.y {
	// 		aCount := diff.x / stepA.x
	// 		return aCount*costA + bCount*costB
	// 	}
	// }
	hcd := HCD(stepB.x, stepB.y)
	x := stepB.x / hcd
	y := stepB.y / hcd
	countA := 0
	countB := 0
	// if (destination.x*y-destination.y*x)%(stepA.x*y-x*stepA.y) == 0 {
	// 	countA = (destination.x*y - destination.y*x) / (stepA.x*y - x*stepA.y)
	// 	countB = (destination.x - stepA.x*countA) / stepB.x
	// 	return countA*costA + countB*costB
	// } else {
	// 	return 0
	// }
	countA = (destination.x*y - destination.y*x) / (stepA.x*y - x*stepA.y)
	countB = (destination.x - stepA.x*countA) / stepB.x

	if countA*stepA.x+countB*stepB.x == destination.x && countA*stepA.y+countB*stepB.y == destination.y {
		return countA*costA + countB*costB
	} else {
		return 0
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error opening file", err)
	}

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		lineA := scanner.Text()
		scanner.Scan()
		lineB := scanner.Text()
		scanner.Scan()
		line := scanner.Text()
		scanner.Scan()

		xa, ya := xy(lineA)
		xb, yb := xy(lineB)
		x, y := xy(line)
		x = x + 10000000000000
		y = y + 10000000000000

		// fmt.Println(xa, ya, xb, yb, x, y)
		destination := point{x, y}
		stepA := point{xa, ya}
		stepB := point{xb, yb}
		costA := 3
		costB := 1
		minCost := minCostOfReaching(destination, stepA, stepB, costA, costB)
		total += minCost
	}
	fmt.Println(total)
}
