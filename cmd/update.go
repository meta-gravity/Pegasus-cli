// // package cmd

// import (
// 	"fmt"
// 	"runtime"
// 	"os/exec"
// 	"os"

// 	"github.com/spf13/cobra"
// )

// // updateCmd represents the update command
// var updateCmd = &cobra.Command{
// 	Use:   "update",
// 	Short: "Update Pegasus CLI to the latest version",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		fmt.Println("Checking for updates...")

// 		// Call the updateCLI function here
// 		if runtime.GOOS == "windows" {
// 			cmd := exec.Command("cmd", "/C", "npm install -g pegasus-cli")
// 			cmd.Stdout = os.Stdout
// 			cmd.Stderr = os.Stderr
// 			if err := cmd.Run(); err != nil {
// 				fmt.Println("Error updating Pegasus CLI:", err)
// 			} else {
// 				fmt.Println("Pegasus CLI has been updated to the latest version!")
// 			}
// 		} else {
// 			cmd := exec.Command("npm", "install", "-g", "pegasus-cli")
// 			cmd.Stdout = os.Stdout
// 			cmd.Stderr = os.Stderr
// 			if err := cmd.Run(); err != nil {
// 				fmt.Println("Error updating Pegasus CLI:", err)
// 			} else {
// 				fmt.Println("Pegasus CLI has been updated to the latest version!")
// 			}
// 		}
// 	},
// }

// func init() {
// 	rootCmd.AddCommand(updateCmd)

// 	// Add a shortcut for the update command
// 	updateCmd.SetAlias("u")
// }