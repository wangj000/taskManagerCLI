/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: `Append item
		- Single Item: task add [item]
		- Multiple Item: task add [item], [item]....`,
	Long: `Add item to currnet list`,
	Run: func(cmd *cobra.Command, args []string) {
		
		data_processed := strings.Join(args, "")
		data := strings.Split(data_processed, ",")

		// Create file if doesn't exist
		if _, err := os.Stat("./test.txt"); err != nil {
			
			// Create txt for task storage
			os.Create("./test.txt")
			
			// Open file for change
			file, _ := os.OpenFile("./test.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
			
			// Add task interface header
			file.WriteString(fmt.Sprintf(`
	╔════════════════════╗
	║    TASK MANAGER    ║
	╚════════════════════╝
══════════════════════════
			`))
		
			// Close file
			file.Close()

		}

		// Open file
		file, err := os.OpenFile("./test.txt", os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
		
		// Close after func
		defer file.Close()

		// Error handling for file creation
		if err != nil {
			fmt.Println("Failed to create file")
			return
		}
		
		// Write todo's to file
		for _, value := range data{
			_, err := file.WriteString(fmt.Sprintf("- %v \n", value)) 
			
			if err != nil {
				fmt.Println("Something went wrong please try again.")
				return
			}

		}

		fmt.Println("Task(s) added")
		return 

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
