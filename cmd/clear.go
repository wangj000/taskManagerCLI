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
	Short: "Clear all tasks",
	Long: `Clear all tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		
		path := filepath.Join("internal", "todos.csv")
		err := os.Remove(path)	

		if err != nil{
			fmt.Println("Failed to clear tasks, please try again")
			return
		}

		fmt.Println("Tasks cleared")
		return

	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clearCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
