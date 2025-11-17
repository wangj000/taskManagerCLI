package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"os"
	// "io"
	"encoding/csv"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Initializes the new TUI cycle instance
		checkTUI := tea.NewProgram(ui.CheckTUI())	

		// finalModel is the last instance of the model 
		// struct after cycle terminates
		finalModel, err := checkTUI.Run()
		if err != nil{
			fmt.Println("Something went wrong with the TUI")
			return
		}
		
		m := finalModel.(ui.CheckModel)

		// Build the path to the resource
		path := filepath.Join("internal", "todos.csv")

		if _, err := os.Stat(path); err != nil {
			fmt.Println("There are no items to to check off")
			return
		}
		
		// Opens the CSV for reading
		file, err := os.Open(path)	
		if err != nil {
			fmt.Println("Something went wrong opening the file")
		}
			
		// Creates the reader, and reads the contents of the CSV 
		reader := csv.NewReader(file) 
		data, err := reader.ReadAll()
		if err != nil{
			fmt.Println("Something went wrong reading the contents of the find chk-64")
			return
		}

		// Changes status of the task item 
		for _, row := range data {
			if row[0] == m.Answers{
				row[len(row) - 1] = "true"
				break
			}
		} 

		// Opens the file for editing
		file, err = os.OpenFile(path, os.O_TRUNC | os.O_WRONLY, 0644)
		if err != nil{
			fmt.Println("Something went wrong opening the file, please try again.")
			return
		}
		defer file.Close()

		// Create and writes to the file
		writer := csv.NewWriter(file)
		err = writer.WriteAll(data)		
		if err != nil{
			fmt.Println("Something went wrong writing to the file")
		}

		return

	},
}

func init() {

	rootCmd.AddCommand(checkCmd)

}
