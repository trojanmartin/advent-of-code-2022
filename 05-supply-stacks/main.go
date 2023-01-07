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

type Port struct {
	stacks map[int]*Stack
}

func (p *Port) move(sourceKey int, destKey int, count int) {
	source, _ := p.stacks[sourceKey]
	dest, _ := p.stacks[destKey]

	for i := 0; i < count; i++ {
		dest.Push(source.Pop())
	}
}

func main() {
	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	port := createPort(scanner)

	fmt.Println(port)
	for scanner.Scan() {

	}
}

func createPort(scanner *bufio.Scanner) *Port {
	port := &Port{
		stacks: make(map[int]*Stack),
	}

	portTextLines := []string{}
	for scanner.Scan() {
		text := scanner.Text()

		if len(text) <= 0 {
			break
		}

		portTextLines = append(portTextLines, text)
	}

	stackNames := portTextLines[len(portTextLines)-1]
	stackCount := 0
	for _, name := range strings.Fields(stackNames) {
		number, err := strconv.Atoi(name)
		if err != nil {
			panic("stack name is not int")
		}
		port.stacks[number] = &Stack{}
		stackCount = number
	}

	currentLine := len(portTextLines) - 2
	for currentLine >= 0 {
		line := portTextLines[currentLine]
		supplies := strings.SplitN(line, " ", stackCount)

		for i, supply := range supplies {
			supply = strings.TrimSpace(supply)
			if len(supply) <= 0 {
				continue
			}
			stack := port.stacks[i+1]
			stack.Push(supply)
		}

		currentLine--
	}

	return port
}
