package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const FILE_NAME = "./data.txt"

type game struct {
	rounds []round
}

type round struct {
	left  string
	right string
}

type score map[string]int
type encoding map[string]string

var scoreMapping = score{
	"A": 1,
	"B": 2,
	"C": 3,
}

var encodingMapping = encoding{
	"X": "A",
	"Y": "B",
	"Z": "C",
}

// KAMEN > NOZNICE
// NOZNICE > PAPIER
// PAPIER > KAMEN
var wins = []string{"AB", "BC", "CA"}

func main() {
	game, _ := readFile(FILE_NAME)
	//result := play(game)
	result := playSecondGame(game)
	fmt.Println("Points in the game: ", result)
}

func GetWinnerPoints(left string, right string) int {
	if left == right {
		return 3
	}

	combination := left + right

	for _, win := range wins {
		if combination == win {
			return 6
		}
	}

	return 0
}

func GetPoints(round round) int {
	rightOption := encodingMapping[round.right]
	leftOption := round.left
	defaultPoints := scoreMapping[rightOption]
	winnerPoints := GetWinnerPoints(leftOption, rightOption)
	return defaultPoints + winnerPoints
}

func play(game *game) int {
	myPoints := 0

	for _, r := range game.rounds {
		myPoints += GetPoints(r)
	}

	return myPoints
}

func getRightOption(round round) string {
	if round.right == "Y" {
		return round.left
	}

	if round.right == "Z" {
		for _, w := range wins {
			lossing := string(w[0])
			winning := string(w[1])
			if lossing == round.left {
				return winning
			}
		}
	}

	if round.right == "X" {
		for _, w := range wins {
			lossing := string(w[0])
			winning := string(w[1])
			if winning == round.left {
				return lossing
			}
		}
	}

	panic("Invalid option")
}

func GetSecondGamePoints(round round) int {
	rightOption := getRightOption(round)
	leftOption := round.left
	defaultPoints := scoreMapping[rightOption]
	winnerPoints := GetWinnerPoints(leftOption, rightOption)
	return defaultPoints + winnerPoints
}

func playSecondGame(game *game) int {
	myPoints := 0

	for _, r := range game.rounds {
		myPoints += GetSecondGamePoints(r)
	}

	return myPoints
}

func readFile(fileName string) (*game, error) {
	game := &game{
		rounds: []round{},
	}

	f, err := os.Open(FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		input := scanner.Text()
		fields := strings.Fields(input)
		game.rounds = append(game.rounds, round{fields[0], fields[1]})
	}

	return game, nil
}
