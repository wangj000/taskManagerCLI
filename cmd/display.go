package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "displays current tasks",
	Long: `displays current tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Checks if file exists
		if _, err := os.Stat("./test.txt"); err != nil {
			fmt.Println("No tasks to display.")	
			return
		}
		
		// Reads the file contents
		data, err := os.ReadFile("./test.txt")
		
		// Error handling for read file
		if err != nil {
			fmt.Println("Something went wrong reading the file, please try again.")
			return 
		}

		fmt.Println(string(data))
		return 

	},
}

func init() {
	rootCmd.AddCommand(displayCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// displayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// displayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
