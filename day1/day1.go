package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// read file
	file, e := os.Open("./input.txt")
	if e != nil {
		panic(e)
	}

	// create scanner
	line := bufio.NewScanner(file)

	// create elves array
	elves := []int{0}

	// iterate through lines
	for line.Scan() {
		if line.Text() == "" {
			// add new elf
			elves = append(elves, 0)
		} else {
			// count calories for elf and append
			calories, e := strconv.Atoi(line.Text())
			if e != nil {
				panic(e)
			}
			elves[len(elves) - 1] +=  calories
		}
	}
	// sort in decending order
	sort.Slice(elves, func(i, j int) bool {return elves[i] > elves[j]})
	// sum top 3 elves
	var sum int
	for _, calories := range elves[:3] {
		sum += calories
	}
	// print result
	fmt.Println(sum)
}
