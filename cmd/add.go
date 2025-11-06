/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	// "strings"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add item to current list",
	Long: `Add item to currnet list`,
	Run: func(cmd *cobra.Command, args []string) {
	
		// Create file if doesn't exist
		filePath, err := os.Create("./test.txt")
		
		// Error handling for file creation
		if err != nil {
			fmt.Println("Failed to create file")
			return
		}
		
		defer filePath.Close()
		
		// Add todo items to list
		// data := make([]byte, 0)
		//
		// for _, task := range args{
		// 	data = append(data, []byte(task)...)
		// }
		
		for _, value := range args{
			_, err := filePath.WriteString(fmt.Sprintf("- %v", value)) 

			if err != nil {
				fmt.Println("Failed to add a task")
				return 
			}

		}

		fmt.Println("Task(s) added")

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
