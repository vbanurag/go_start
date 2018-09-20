package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var userInput int
	var secretNumber int
	var quit bool = false

	rand.Seed(time.Now().Unix())
	secretNumber = rand.Intn(10)

	for quit != true {
		fmt.Printf("Please Enter a number: ")
		fmt.Scan(&userInput)
		if userInput == secretNumber {
			fmt.Println("You Won")
			quit = true
		} else if userInput < secretNumber {
			fmt.Println("Too Low...")
		} else if userInput > secretNumber {
			fmt.Println("Too High !!")
		}
	}

}
