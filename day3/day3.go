package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func mulPattern(s string, i int) (int, int) {
	if i+8 <= len(s) {
		if s[i] == 'm' && s[i+1] == 'u' && s[i+2] == 'l' && s[i+3] == '(' {
			endSearchIndex := min(i+11, len(s)-1)

			subs := s[i+4 : endSearchIndex+1]
			endIndex := strings.IndexRune(subs, ')')

			if endIndex == -1 {
				return 0, 0
			}

			subs = s[i+4 : i+4+endIndex+1]

			comaIndex := strings.IndexRune(subs, ',')

			if comaIndex == -1 {
				return 0, 0
			}

			num1, err := strconv.Atoi(s[i+4 : i+4+comaIndex])

			if err != nil {
				return 0, 0
			}

			num2, err := strconv.Atoi(s[i+4+comaIndex+1 : i+4+endIndex])

			if err != nil {
				return 0, 0
			}

			return num1, num2

		}

	}

	return 0, 0
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

	bytes, err := io.ReadAll(file)

	if err != nil {
		log.Fatal("Error opening the file : %w", err)
	}

	text := string(bytes)
	total := 0
	i := 0
	for i < len(text) {
		nextText := text[i:]

		dontIndex := strings.Index(nextText, "don't()")

		if dontIndex == -1 {
			dontIndex = len(nextText) - 1
		}

		subText := nextText[:dontIndex]

		for i := range subText {
			x, y := mulPattern(subText, i)
			total += (x * y)
		}

		doText := nextText[dontIndex:]

		doIndex := strings.Index(doText, "do()")

		if doIndex == -1 {
			break
		}

		i += dontIndex + doIndex
	}

	fmt.Println(total)

}
