package part2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Main() {
	file, e := os.Open("./input.txt")
	if e != nil {
		panic(e)
	}
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		// get next three bags
		var bags []string
		bags = append(bags, scanner.Text())
		for i := 0; i < 2; i++ {
			scanner.Scan()
			bags = append(bags, scanner.Text())
		}
		// find common character between strings
		charIndex := strings.IndexFunc(
			bags[0],
			func (char rune) bool {return strings.ContainsRune(bags[1], char) && strings.ContainsRune(bags[2], char)},
		)
		char := bags[0][charIndex]
		// append to sum
		sum += itemPriority(char)
	}
	fmt.Println(sum)
}

func itemPriority(char byte) int {
	if (char <= 'Z' && char >= 'A') {
		return int(char) - 38
	} else if (char <= 'z' && char >= 'a') {
		return int(char) - 96
	}
	return 0
}

