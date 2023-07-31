package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	// "golang.org/x/text/cases"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("GravityX CLI Program!")
	fmt.Println("---------------------------------")

	for {
		fmt.Print("Enter a command ('exit' or 'leave' to quit):")
		fmt.Print("\nEnter a command ('clear' to clear): ")
		input, _ := reader.ReadString('\n')

		// Trim any leading/trailing whitespaces and newlines
		input = strings.TrimSpace(input) 
		if input == "clear" {
			fmt.Println("Clearing the program...please wait...")
			break
		}

		if input == "exit" {
			fmt.Println("Exiting the program...")
			break
		}
		if input == "leave" {
			fmt.Println("Exiting the program...")
			break
		}
		// Process the command
		result := processCommand(input)
		fmt.Println(result)
	}

	fmt.Println("Thank you for using the Simple CLI Program!")
}

func processCommand(command string) string {
	// Add your logic to process different commands here
	switch command {
	case "hi":
		return "Hello, there!"
	case "Author":
		return "Metagravity"
	case "time":
		return "The current time is " + time.Now().Format("15:04:05")

	case "month":
		return "The current month is " + time.Now().Month().String()

	case "year":
		return "The current year is " + fmt.Sprintf("%d", time.Now().Local().Year())

	case "color red":
		color.Set(color.FgRed)
		defer color.Unset()
		return "The color is set to red."
	case "color green":
		color.Set(color.FgGreen)
		defer color.Unset()
		return "The color is set to green."
	case "color blue":
		color.Set(color.FgBlue)
		defer color.Unset()
		return "The color is set to blue."
	default:
		return "Unknown command. Please try again."
	}


}

func myTime() {
	// currentTime := time.Now()
	time.Now().Month()
	time.Now().Local().Year()
}

// func continue() {
	
// }