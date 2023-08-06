// package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/spf13/cobra"
// )

// var divideCmd = &cobra.Command{
// 	Use:   "divide",
// 	Short: "Divide numbers",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) < 2 {
// 			fmt.Println("Please provide at least two numbers to divide")
// 			return
// 		}

// 		// Convert arguments to integers
// 		num1, err := strconv.Atoi(args[0])
// 		if err != nil {
// 			fmt.Printf("Invalid number: %s\n", args[0])
// 			return
// 		}

// 		num2, err := strconv.Atoi(args[1])
// 		if err != nil {
// 			fmt.Printf("Invalid number: %s\n", args[1])
// 			return
// 		}

// 		// Check for division by zero
// 		if num2 == 0 {
// 			fmt.Println("Cannot divide by zero")
// 			return
// 		}

// 		// Perform the division
// 		result := num1 / num2
// 		fmt.Printf("Result: %d\n", result)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(divideCmd)
// }