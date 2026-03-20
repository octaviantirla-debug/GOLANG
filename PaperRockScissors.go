package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("🌱 Welcome to Paper, Rock, Scissors! 🌱")
	fmt.Println("Choose: paper, rock, or scissors (or 'quit' to exit)")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	for {
		fmt.Print("Your choice: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		userInput = strings.ToLower(userInput)

		// Exit condition
		if userInput == "quit" {
			fmt.Println("Thanks for playing! Goodbye. 🌿")
			break
		}

		// Validate user input
		if userInput != "paper" && userInput != "rock" && userInput != "scissors" {
			fmt.Println("❌ Invalid choice. Please enter 'paper', 'rock', 'scissors', or 'quit'.")
			continue
		}

		// Computer makes a random choice
		choices := []string{"paper", "rock", "scissors"}
		computerChoice := choices[rand.Intn(len(choices))]
		fmt.Printf("Computer chose: %s\n", computerChoice)

		// Determine the winner
		if userInput == computerChoice {
			fmt.Println("🤝 It's a tie!")
		} else if (userInput == "paper" && computerChoice == "rock") ||
			(userInput == "rock" && computerChoice == "scissors") ||
			(userInput == "scissors" && computerChoice == "paper") {
			fmt.Println("🎉 You win!")
		} else {
			fmt.Println("😢 Computer wins!")
		}
	}
}
