package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const FILE_NAME = "./data.txt"

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fullOverlapCount := 0
	overlapCount := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		segments := strings.Split(line, ",")

		left := segments[0]
		right := segments[1]
		leftParsed := strings.Split(left, "-")
		leftStart, _ := strconv.Atoi(leftParsed[0])
		leftEnd, _ := strconv.Atoi(leftParsed[1])

		rightParsed := strings.Split(right, "-")
		rightStart, _ := strconv.Atoi(rightParsed[0])
		rightEnd, _ := strconv.Atoi(rightParsed[1])

		if isFullOverlap(leftStart, leftEnd, rightStart, rightEnd) {
			fullOverlapCount++
		}

		if isOverlap(leftStart, leftEnd, rightStart, rightEnd) {
			overlapCount++
		}
	}

	fmt.Println("Sum of full overlaps is: ", fullOverlapCount)
	fmt.Println("Sum of partial overlaps is: ", overlapCount)
}

func isOverlap(leftStart int, leftEnd int, rightStart int, rightEnd int) bool {
	if leftStart == rightStart {
		return true
	}

	if leftStart < rightStart {
		return rightStart <= leftEnd
	}

	if leftStart > rightStart {
		return rightEnd >= leftStart
	}

	return false
}

func isFullOverlap(leftStart int, leftEnd int, rightStart int, rightEnd int) bool {

	if leftStart == rightStart {
		return true
	}

	if leftStart < rightStart {
		return leftEnd >= rightEnd
	}

	if leftStart > rightStart {
		return leftEnd <= rightEnd
	}

	return false
}
