package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	

	"github.com/fatih/color"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Simple CLI Program!")
	fmt.Println("---------------------------------")

	for {
		fmt.Print("Enter a command ('exit' or 'leave' to quit): ")
		input, _ := reader.ReadString('\n')

		// Trim any leading/trailing whitespaces and newlines
		input = strings.TrimSpace(input)

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
	case "hello":
		return "Hello, there!"
	case "time":
		return "The current time is " + getCurrentTime()
	case "jazz":
		return "this is meta jazz and the time is" + getCurrentTime()
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

func getCurrentTime() string {
	return time.Now().Format("15:04:05")
}
