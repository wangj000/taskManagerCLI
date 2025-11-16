package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"path/filepath"
	"os"
	"io"
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

		// General workflow:
		// - Get user input of the item that they want to check 
		// - Read from the file directly and change the row (ID - 1) index
		// - Change the status to True
	
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
		}

		// Open/Close the file
		file, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY | os.O_TRUNC, 0644)	
		if err != nil{
			fmt.Println("Something went wrong opening the file, please try again.")
			return
		}
		defer file.Close()
		
		// Creates a reader in the file that we opened
		reader := csv.NewReader(file) 
		data := make([][]string, 0)

		for {

			record, err := reader.Read()

			if err == io.EOF{
				break
			}

			// FIX: For some reason it's erroring out right here when i try to read the file
			if err != nil {
				fmt.Printf("%v Something went wrong reading the file, please try again", err)
				return
			}
			
			if m.Answers == record[0]{
				record[len(record) - 1] = "true"
			}

			data = append(data, record)

		}

	
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
