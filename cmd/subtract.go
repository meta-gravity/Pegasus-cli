// package main

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/spf13/cobra"
// )

// var subtractCmd = &cobra.Command{
// 	Use:   "subtract",
// 	Short: "Subtract numbers",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		result := 0
// 		for _, arg := range args {
// 			num, err := strconv.Atoi(arg)
// 			if err != nil {
// 				fmt.Printf("Invalid number: %s\n", arg)
// 				return
// 			}
// 			result -= num
// 		}
// 		fmt.Printf("Result: %d\n", result)
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(subtractCmd)
// }