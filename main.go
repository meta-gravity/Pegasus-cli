package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
	// "runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const version = "1.0.0"
const githuLink = "https://github.com/meta-gravity/Pegasus-cli"
const link = "https://github.com/meta-gravity/Pegasus-cli/releases/latest"
const asciiArt = `
_____           _____           _____                    _____                    _____           _____                   
/\    \         /\    \         /\    \                  /\    \                  /\    \         /\    \                  
/::\____\       /::\    \       /::\    \                /::\    \                /::\    \       /::\    \                 
/:::/    /      /::::\    \     /::::\    \              /::::\    \              /::::\    \     /::::\    \                
/:::/    /      /::::::\    \   /::::::\    \            /::::::\    \            /::::::\    \   /::::::\    \               
/:::/    /      /:::/\:::\    \ /:::/\:::\    \          /:::/\:::\    \          /:::/\:::\    \ /:::/\:::\    \              
/:::/____/      /:::/__\:::\    /:::/__\:::\    \        /:::/__\:::\    \        /:::/  \:::\    /:::/__\:::\    \             
\:::\    \     /::::\   \:::\  /::::\   \:::\    \      /::::\   \:::\    \      /:::/    \:::\  /::::\   \:::\    \            
\:::\    \   /::::::\   \:::\/::::::\   \:::\    \    /::::::\   \:::\    \    /:::/    / \:::\/::::::\   \:::\    \           
\:::\    \ /:::/\:::\   \:::::::/\:::\   \:::\    \  /:::/\:::\   \:::\    \  /:::/    /   \:::::::/\:::\   \:::\    \          
\:::\    /:::/  \:::\   \::::/  \:::\   \:::\    \/:::/  \:::\   \:::\____\/:::/____/     \::::/  \:::\   \:::\____\         
 \:::\  /:::/    \:::\  /:::/    \:::\  /:::\    \/:::/    \:::\   \::/    /\:::\    \     /:::/    \:::\   \::/    /         
  \:::\  /:::/    / \:::\/:::/    / \:::\/:::/    / \::/    / \:::\   \/____/  \:::\    \   /:::/    / \:::\   \/____/          
   \:::\/:::/    /   \::::::/    /   \::::::/    /   \/____/   \:::\    \       \:::\    \ /:::/    /   \:::\    \              
	\::::::/    /     \::::/____/     \::::/____/              \:::\____\       \:::\    /:::/    /     \:::\____\             
	 \::::/    /       \:::\    \      \:::\    \               \::/    /        \:::\  /:::/    /       \::/    /             
	  \::/____/         \:::\    \      \:::\    \              /::/    /          \:::\/:::/    /         \/____/              
	   ~~                \:::\    \      \:::\    \            /::/    /            \::::::/    /                              
						  \:::\____\      \:::\____\          /::/    /              \::::/    /                               
						   \::/    /       \::/    /          \::/    /                \::/____/                                
							\/____/         \/____/            \/____/                                                           
`



var (
	currentColor color.Attribute // Holds the current color setting
	history      []string         // Slice to store command history
	rootCmd      *cobra.Command   // Declare rootCmd as a global variable
)


func main() {
	fmt.Println("Pegasus CLI")
	fmt.Println("---------------------------------------------------------------------------------------------------------------")
    printWelcomeMessage()
	color.Blue(asciiArt)

	rootCmd = &cobra.Command{Use: "Pegasus"}





	// Add commands to the root command
	rootCmd.AddCommand(hiCmd)
	rootCmd.AddCommand(authorCmd)
	rootCmd.AddCommand(currentTimeCmd)
	rootCmd.AddCommand(currentMonthCmd)
	rootCmd.AddCommand(currentYearCmd)
	rootCmd.AddCommand(colorCmd)
	rootCmd.AddCommand(clearCmd)
	// rootCmd.AddCommand(historyCmd)
	rootCmd.AddCommand(versionCmd)
	// rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(linkCmd)
	// for updates well still in progress
	// rootCmd.AddCommand(updateCmd)

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// After the command execution, enter the interactive mode
	interactiveMode()
}

// var updateCmd = &cobra.Command{
//     Use:   "update",
//     Short: "Update Pegasus CLI to the latest version",
//     Run: func(cmd *cobra.Command, args []string) {
//         err := updateCLI()
//         if err != nil {
//             fmt.Println("Update failed:", err)
//         }
//     },
// }

func showList() string {
	listText := "Available Commands:\n"
	for _, c := range rootCmd.Commands() {
		listText += fmt.Sprintf("- %s: %s\n", c.Use, c.Short)
	}
	return listText
}

func printWelcomeMessage() {
    fmt.Printf(asciiArt)
}


func interactiveMode() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a command ('exit' or 'leave' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" || input == "leave" || input == "quit" {
			fmt.Println("Thank you for using the Pegasus CLI Program!")
			break
		} else if input == "clear" || input == "cls" {
			clearScreen()
		} else {
			// Process the command
			result := processCommand(input)
			fmt.Println(result)
			// Add a history entry
			history = append(history, input)
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
	case "color":
		return colorString("color red\ncolor blue\ncolor cyan\ncolor white\ncolor magenta\ncolor green", currentColor)
	case "color red":
		currentColor = color.FgHiRed
		return colorString("The color is set to red.", currentColor)
	case "color white":
		currentColor = color.FgHiWhite
		return colorString("The color has been set to default.", currentColor)
	case "color cyan":
		currentColor = color.FgHiCyan
		return colorString("The color is set to cyan.", currentColor)
	case "color magenta":
		currentColor = color.FgHiMagenta
		return colorString("The color is set to magenta.", currentColor)
	case "color green":
		currentColor = color.FgHiGreen
		return colorString("The color is set to green.", currentColor)
	case "color blue":
		currentColor = color.FgHiBlue
		return colorString("The color is set to blue.", currentColor)
	case "history":
		return showHistory()
	case "version":
		return showVersion()
	case "list":
		return showList()
	case "github":
		return showgithubLink()
	case "file":
		return "file" // Placeholder for future implementation
	// case "update":
	// 	return showUpdate()
	default:
		return color.HiRedString("Unknown command.")
	}
}

func showHistory() string {
	historyText := "Command History:\n"
	for i, cmd := range history {
		historyText += fmt.Sprintf("%d. %s\n", i+1, cmd)
	}
	return historyText
}

func showgithubLink() string {
	return fmt.Sprintf("Pegasus CLI github link: %s", link)
}

func showVersion() string {
	return fmt.Sprintf("Pegasus CLI Version: %s", version)
}

func colorString(s string, c color.Attribute) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", c, s)
}

func clearScreen() {
	cmd := exec.Command("clear") // Use "cls" on Windows
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("Pegasus CLI")
	fmt.Println("---------------------------------")
}

// func showUpdate(cmd *cobra.Command, args []string) {
// 	// Implement the update logic here
// 	fmt.Println("Updating Pegasus CLI...")
// 	// Add the actual update logic here
// }

// still working on this part 
// func showUpdate(cmd *cobra.Command, args []string) {
// 	// Step 1: Retrieve the latest release version from GitHub
// 	latestVersion, err := getLatestReleaseVersion()
// 	if err != nil {
// 		fmt.Println("Failed to retrieve the latest release version:", err)
// 		return
// 	}

// 	// Step 2: Compare the latest version with the current version
// 	if latestVersion == version {
// 		fmt.Println("Your CLI is already up to date.")
// 		return
// 	}

// 	// Step 3: Download the new release binary from GitHub
// 	downloadURL := fmt.Sprintf("https://github.com/YOUR_USERNAME/YOUR_REPOSITORY/releases/download/%s/pegasus.exe", latestVersion)
// 	resp, err := http.Get(downloadURL)
// 	if err != nil {
// 		fmt.Println("Failed to download the update:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	// Step 4: Replace the old binary with the new one
// 	updatePath := "bin/pegasus.exe" // The path where the new binary will be saved
// 	out, err := os.Create(updatePath)
// 	if err != nil {
// 		fmt.Println("Failed to create the update file:", err)
// 		return
// 	}
// 	defer out.Close()

// 	_, err = io.Copy(out, resp.Body)
// 	if err != nil {
// 		fmt.Println("Failed to save the update file:", err)
// 		return
// 	}

// 	fmt.Println("Pegasus CLI has been updated to the latest version.")
// }

var readCmd = &cobra.Command{
	Use:   "read <file>",
	Short: "Read the contents of a file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]
		data, err := ioutil.ReadFile(file)
		if err != nil {
			// Check if the file does not exist
			if os.IsNotExist(err) {
				fmt.Printf("Error: File '%s' does not exist.\n", file)
			} else {
				fmt.Printf("Error reading file '%s': %v\n", file, err)
			}
			return
		}
		fmt.Println(string(data))
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

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pegasus CLI Version:", version)
	},
}

var linkCmd = &cobra.Command{
	Use: "link",
	Short: "contribute to this project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Pegasus CLI github link:", link)
	},
}

var colorCmd = &cobra.Command{
	Use:   "color <color>",
	Short: "Sets the color",
	Args:  cobra.ExactArgs(1), // Expects exactly one argument after "color" command
	Run: func(cmd *cobra.Command, args []string) {
		colorString := strings.TrimSpace(args[0])
		colorAttr, err := getColorAttribute(colorString)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		currentColor = colorAttr
		fmt.Printf("The color is set to %s.\n", colorString)
	},
}

func getColorAttribute(colorString string) (color.Attribute, error) {
	switch strings.ToLower(colorString) {
	case "red":
		return color.FgRed, nil
	case "green":
		return color.FgGreen, nil
	case "blue":
		return color.FgBlue, nil
	case "cyan":
		return color.FgCyan, nil
	case "white":
		return color.FgWhite, nil
	case "magenta":
		return color.FgMagenta, nil
	default:
		return 0, fmt.Errorf("invalid color: %s. Available colors: red, green, blue, cyan, white, magenta", colorString)
	}
}

