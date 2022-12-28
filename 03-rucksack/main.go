package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const FILE_NAME = "./data.txt"

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	// fmt.Println("Sum of the double items priorities is: ", doubleItemsPriority(f))

	fmt.Println("Sum of badges item types priorities: ", scanBadges(f))
}

func doubleItemsPriority(f *os.File) int {
	priority := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		priority += scanRucksackForDoubleItems(scanner.Text())
	}
	return priority
}

func scanBadges(f *os.File) int {
	priority := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		r1 := scanner.Text()
		scanner.Scan()
		r2 := scanner.Text()
		scanner.Scan()
		r3 := scanner.Text()
		priority += scanRucksackForBadges(r1, r2, r3)
	}
	return priority
}

func scanRucksackForBadges(r1 string, r2 string, r3 string) int {
	occurencies := make(map[int]int)

	for _, item := range r1 {
		occurencies[int(item)] = 1
	}

	for _, item := range r2 {
		_, exist := occurencies[int(item)]

		if exist {
			occurencies[int(item)] = 2
		}
	}

	for _, item := range r3 {
		occ, exist := occurencies[int(item)]
		if exist && occ == 2 {
			return getItemPriority(int(item))
		}
	}

	fmt.Println(occurencies)

	panic("No badge found")
}

func scanRucksackForDoubleItems(rucksack string) int {
	midPoint := len(rucksack) / 2

	left := rucksack[:midPoint]
	right := rucksack[midPoint:]

	occurencies := make(map[int]bool, midPoint)

	for _, item := range left {
		occurencies[int(item)] = true
	}

	priority := 0
	for _, item := range right {
		if occurencies[int(item)] {
			priority += getItemPriority(int(item))
			occurencies[int(item)] = false
		}
	}
	return priority
}

func getItemPriority(item int) int {
	if 'a' <= item && item <= 'z' {
		return item - 96
	}

	if 'A' <= item && item <= 'Z' {
		return item - 38
	}

	panic("Invalid item")
}
