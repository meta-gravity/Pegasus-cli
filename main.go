package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
	"time"
)

var currentColor color.Attribute // Holds the current color setting

func main() {
	fmt.Println("Pegasus cli")
	fmt.Println("---------------------------------")

	rootCmd := &cobra.Command{Use: "mycli"}

	// Add commands to the root command
	rootCmd.AddCommand(hiCmd)
	rootCmd.AddCommand(authorCmd)
	rootCmd.AddCommand(currentTimeCmd)
	rootCmd.AddCommand(currentMonthCmd)
	rootCmd.AddCommand(currentYearCmd)
	rootCmd.AddCommand(colorCmd)
	rootCmd.AddCommand(clearCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// After the command execution, enter the interactive mode
	interactiveMode()
}

func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a command ('exit' or 'leave' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "leave" {
			fmt.Println("Thank you for using the Simple CLI Program!")
			break
		} else if input == "clear" {
			clearScreen()
		} else {
			// Process the command
			result := processCommand(input)
			fmt.Println(result)
		}
	}
}

func processCommand(command string) string {
	switch command {
	case "hello":
		return colorString("Hello, there! 👋", currentColor)
	case "hi":
		return colorString("Hi! 😊", currentColor)
	case "band":
		return colorString("Hello from the band! 🎸🥁", currentColor)
	case "halo":
		return colorString("Halo there! 🌞", currentColor)
	case "author":
		return colorString("Metagravity", currentColor)
	case "time":
		return colorString("The current time is "+time.Now().Format("15:04:05"), currentColor)
	case "month":
		return colorString("The current month is "+time.Now().Month().String(), currentColor)
	case "year":
		return colorString("The current year is "+fmt.Sprintf("%d", time.Now().Local().Year()), currentColor)
	case "color red":
		currentColor = color.FgRed
		return colorString("The color is set to red.", currentColor)
	case "color green":
		currentColor = color.FgGreen
		return colorString("The color is set to green.", currentColor)
	case "color blue":
		currentColor = color.FgBlue
		return colorString("The color is set to blue.", currentColor)
	default:
		return "Unknown command. Please try again."
	}
}

func colorString(s string, c color.Attribute) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("Pegasus cli")
	fmt.Println("---------------------------------")
}

// Rest of the code remains the same

var hiCmd = &cobra.Command{
	Use:   "hi",
	Short: "Prints a greeting message",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorString("Hello, there!", currentColor))
	},
}

var authorCmd = &cobra.Command{
	Use:   "author",
	Short: "Prints the author's name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorString("Metagravity", currentColor))
	},
}

var currentTimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Prints the current time",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorString("The current time is "+time.Now().Format("15:04:05"), currentColor))
	},
}

var currentMonthCmd = &cobra.Command{
	Use:   "month",
	Short: "Prints the current month",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorString("The current month is "+time.Now().Month().String(), currentColor))
	},
}

var currentYearCmd = &cobra.Command{
	Use:   "year",
	Short: "Prints the current year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(colorString("The current year is "+fmt.Sprintf("%d", time.Now().Local().Year()), currentColor))
	},
}

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the screen",
	Run: func(cmd *cobra.Command, args []string) {
		clearScreen()
	},
}

var colorCmd = &cobra.Command{
	Use:   "color <color>",
	Short: "Sets the color",
	Args:  cobra.ExactArgs(1), // Expects exactly one argument after "color" command
	Run: func(cmd *cobra.Command, args []string) {
		colorString := strings.TrimSpace(args[0])
		switch colorString {
		case "red":
			currentColor = color.FgRed
		case "green":
			currentColor = color.FgGreen
		case "blue":
			currentColor = color.FgBlue
		default:
			fmt.Println("Invalid color. Available colors: red, green, blue")
			return
		}
		fmt.Printf("The color is set to %s.\n", colorString)
	},
}

// Commands to set the color using aliases (shortcuts)
var colorRedCmd = &cobra.Command{
	Use:    "red",
	Hidden: true, // Hide the command from help display
	Run: func(cmd *cobra.Command, args []string) {
		colorCmd.Run(cmd, []string{"red"})
	},
}

var colorGreenCmd = &cobra.Command{
	Use:    "green",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		colorCmd.Run(cmd, []string{"green"})
	},
}

var colorBlueCmd = &cobra.Command{
	Use:    "blue",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		colorCmd.Run(cmd, []string{"blue"})
	},
}

func init() {
	// Add aliases for color commands
	colorCmd.AddCommand(colorRedCmd)
	colorCmd.AddCommand(colorGreenCmd)
	colorCmd.AddCommand(colorBlueCmd)
}