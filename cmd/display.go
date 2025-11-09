package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"io"
	"encoding/csv"
	"path/filepath"	
)

var displayCmd = &cobra.Command{
	Use:   "display",
	Short: "displays current tasks",
	Long: `displays current tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		
		path := filepath.Join("internal", "todos.csv")

		// Checks if file exists
		if _, err := os.Stat(path); err != nil {
			fmt.Println("No tasks to display, try adding a task.")	
			return
		}
		
		// Opening the file
		file, err := os.Open(path)
		if err != nil{
			fmt.Println("Something went wrong opening the file, please try again.")
			return
		}
		defer file.Close()
		
		// BUG: I think the bug is triggering here, because something is going on with the reader
		reader := csv.NewReader(file)
		for {

			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("%v second print - Something went wrong reading the file, please try again", err)
				return
			}
			
			for _, item := range record {
				fmt.Printf(item + " ")
			}

			fmt.Println("\n")

		}
		
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
