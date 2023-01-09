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

type Port struct {
	stacks map[int]*Stack
}

func (p *Port) move(sourceKey int, destKey int, count int) {
	source := p.stacks[sourceKey]
	dest := p.stacks[destKey]

	for i := 0; i < count; i++ {
		dest.Push(*source.Pop())
	}
}

func (p *Port) movePart2(sourceKey int, destKey int, count int) {
	source := p.stacks[sourceKey]
	dest := p.stacks[destKey]

	toMove := []ItemType{}

	for i := 0; i < count; i++ {
		toMove = append(toMove, *source.Pop())
	}

	for i := count - 1; i >= 0; i-- {
		dest.Push(toMove[i])
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
		fields := strings.Fields(scanner.Text())
		count, _ := strconv.Atoi(fields[1])
		from, _ := strconv.Atoi(fields[3])
		to, _ := strconv.Atoi(fields[5])

		port.movePart2(from, to, count)
	}

	for i := 1; i <= 9; i++ {
		stack := port.stacks[i]
		fmt.Print(*stack.Pop())
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

		i := 0
		stackIndex := 1
		for stackIndex <= stackCount {
			end := i + 4
			if end > len(line) {
				end = len(line)
			}
			supply := line[i:end]

			supply = strings.TrimSpace(supply)
			if len(supply) > 0 {
				stack := port.stacks[stackIndex]
				stack.Push(supply)
			}
			stackIndex++
			i += 4
		}

		currentLine--
	}

	return port
}
