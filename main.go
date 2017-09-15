package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var resultText string
	var standing = []int{0, 0}
	var figures = []string{"Rock", "Paper", "Scissors", "Exit"}

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

			switch playerChoiceVerbose {
			case aiChoiceVerbose:
				resultText = "Your " + figures[choice] + " draws with AI's " + figures[aiChoice] + "!"
			case "Rock":
				if aiChoiceVerbose == "Scissors" {
					resultText = "Your " + figures[choice] + " beats AI's " + figures[aiChoice] + "!"
					standing[0]++
				} else {
					resultText = "Your " + figures[choice] + " beaten by AI's " + figures[aiChoice] + "!"
					standing[1]++
				}
			case "Paper":
				if aiChoiceVerbose == "Rock" {
					resultText = "Your " + figures[choice] + " beats AI's " + figures[aiChoice] + "!"
					standing[0]++
				} else {
					resultText = "Your " + figures[choice] + " beaten by AI's " + figures[aiChoice] + "!"
					standing[1]++
				}
			case "Scissors":
				if aiChoiceVerbose == "Paper" {
					resultText = "Your " + figures[choice] + " beats AI's " + figures[aiChoice] + "!"
					standing[0]++
				} else {
					resultText = "Your " + figures[choice] + " beaten by AI's " + figures[aiChoice] + "!"
					standing[1]++
				}
			}

			loop++
		}
	}
}
