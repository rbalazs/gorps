package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type GameHandler interface {
	Win(string, string) (string, int)
	Lose(string, string) (string, int)
}

type RockPaperScissors struct {
}

func (rps RockPaperScissors) Win(aiChoiceVerbose string, playerChoiceVerbose string) (string, int) {
	return "Your " + playerChoiceVerbose + " beats AI's " + aiChoiceVerbose + "!", 0
}

func (rps RockPaperScissors) Lose(aiChoiceVerbose string, playerChoiceVerbose string) (string, int) {
	return "Your " + playerChoiceVerbose + " beaten by AI's " + aiChoiceVerbose + "!", 1
}

func main() {
	var resultText string
	var standing = []int{0, 0}
	var figures = []string{"Rock", "Paper", "Scissors", "Exit"}
	var gameHandler = RockPaperScissors{}

	exit := false
	loop := 0

	for exit == false {
		var choice int
		var aiChoice int

		fmt.Printf("\033[2J")
		fmt.Printf("\033[H")

		if loop == 0 {
			fmt.Println("\n\nYo. I want to play a game!")
			fmt.Println("\n\n\n\tPick one:")
		} else {
			fmt.Println("\n\n" + resultText)
			fmt.Println("\nThe standing is: " + strconv.Itoa(standing[0]) + " to " + strconv.Itoa(standing[1]) + ".")
			fmt.Println("\n\tOne more time?")
		}

		for index, element := range figures {
			fmt.Println(map[bool]string{true: "\n", false: ""}[index == 0] + strconv.Itoa(index) + ".\t" + element)
		}

		fmt.Scanf("%d", &choice)

		if choice == 3 {
			fmt.Println("\t\tBye!")
			exit = true
		} else {
			aiChoice = rand.Intn(3)
			aiChoiceVerbose := figures[aiChoice]
			playerChoiceVerbose := figures[choice]
			standingIndexToInc := -1

			switch playerChoiceVerbose {
			case aiChoiceVerbose:
				resultText = "Your " + figures[choice] + " draws with AI's " + figures[aiChoice] + "!"
				continue
			case "Rock":
				if aiChoiceVerbose == "Scissors" {
					resultText, standingIndexToInc = gameHandler.Win(aiChoiceVerbose, playerChoiceVerbose)
				} else {
					resultText, standingIndexToInc = gameHandler.Lose(aiChoiceVerbose, playerChoiceVerbose)
				}
			case "Paper":
				if aiChoiceVerbose == "Rock" {
					resultText, standingIndexToInc = gameHandler.Win(aiChoiceVerbose, playerChoiceVerbose)
				} else {
					resultText, standingIndexToInc = gameHandler.Lose(aiChoiceVerbose, playerChoiceVerbose)
				}
			case "Scissors":
				if aiChoiceVerbose == "Paper" {
					resultText, standingIndexToInc = gameHandler.Win(aiChoiceVerbose, playerChoiceVerbose)
				} else {
					resultText, standingIndexToInc = gameHandler.Lose(aiChoiceVerbose, playerChoiceVerbose)
				}
			}

			standing[standingIndexToInc]++
			loop++
		}
	}
}
