package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FILE_NAME = "./data.txt"

func main() {
	log.SetFlags(0)
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer closeFile(f)

	scanner := bufio.NewScanner(f)

	maxTotalCalories := 0
	currentCalories := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if currentCalories > maxTotalCalories {
				maxTotalCalories = currentCalories
			}
			currentCalories = 0
			continue
		}

		lineCal, err := strconv.Atoi(line)
		checkError(err)
		currentCalories += lineCal
	}

	if currentCalories > maxTotalCalories {
		maxTotalCalories = currentCalories
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Max total calories: ", maxTotalCalories)
}

func closeFile(f *os.File) {
	err := f.Close()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
