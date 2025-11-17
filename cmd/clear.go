package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"path/filepath"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "To clear all tasks from the table",
	Long: `To clear all the tasks present in the table, by using
	'task clear' which deletes the CSV that contains all the task
	information`,
	Run: func(cmd *cobra.Command, args []string) {
		
		path := filepath.Join("internal", "todos.csv")
		err := os.Remove(path)	
		if err != nil{
			fmt.println("failed to clear tasks, please try again")
			return
		}


		fmt.Println("Tasks cleared")
		return

	},
}

func init() {
	rootCmd.AddCommand(clearCmd)
}
