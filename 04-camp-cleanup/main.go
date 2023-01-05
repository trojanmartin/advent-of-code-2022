package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "./test-data.txt"

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	overlapCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, ",")

		fmt.Print(line)

		if isOverlap(segments[0], segments[1]) {
			overlapCount++
			fmt.Println("   ", "True")
		}
		fmt.Println()
	}

	fmt.Println("Sum of overlaps is: ", overlapCount)
}

func isOverlap(left string, right string) bool {
	leftParsed := strings.Split(left, "-")
	leftStart, _ := strconv.Atoi(leftParsed[0])
	leftEnd, _ := strconv.Atoi(leftParsed[1])

	rightParsed := strings.Split(right, "-")
	rightStart, _ := strconv.Atoi(rightParsed[0])
	rightEnd, _ := strconv.Atoi(rightParsed[1])

	return false
}
