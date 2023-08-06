// package main

// import (
// 	"fmt"

// 	"github.com/spf13/cobra"
// )

// var addCmd = &cobra.Command{
// 	Use:   "add",
// 	Short: "Add two numbers",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		if len(args) < 2 {
// 			fmt.Println("Usage: Pegasus add <num1> <num2>")
// 			return
// 		}

// 		num1 := parseNumber(args[0])
// 		num2 := parseNumber(args[1])

// 		result := num1 + num2
// 		fmt.Printf("Result: %f\n", result)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(addCmd)
// }

// func parseNumber(s string) float64 {
// 	var num float64
// 	_, err := fmt.Sscanf(s, "%f", &num)
// 	if err != nil {
// 		return 0
// 	}
// 	return num
// }