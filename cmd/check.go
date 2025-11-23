package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"os"
	"encoding/csv"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "To mark items as completed",
	Long: `Used to change the status of items in the list. For example
	if you do 'task check' and then put in the task ID as prompted
	by the TUI you can change the status to 'true'`,
	Run: func(cmd *cobra.Command, args []string) {

		// Intializes new program instance
		checkTUI := tea.NewProgram(ui.CheckTUI())	

		// The latest struct post program termination
		inalModel, err := checkTUI.Run()
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

		// Create + writes to file
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
