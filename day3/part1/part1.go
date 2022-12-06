package part1

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
		bag := scanner.Text()
		// split into two
		compartment1 := bag[:len(bag) / 2]
		compartment2 := bag[len(bag) / 2:]
		// compare for repeat characters
		charIndex := strings.IndexAny(compartment1, compartment2)
		char := compartment1[charIndex]
		// get repeat character priority and add to sum
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
