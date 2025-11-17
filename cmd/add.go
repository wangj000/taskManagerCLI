package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"os"
	"encoding/csv"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
	filetool "github.com/wangj000/task/utils"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: `To add new items to the list`,
	Long: `To append new elements to the list use 
	'task add' when put in the task details as prompted
	by the TUI and upon re-rendering you should see the
	updated table.`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Initializes TUI program    
		addTUI := tea.NewProgram(ui.AddTUI())	

		// The latest struct post process termination
		finalModel, err := addTUI.Run()
		if err != nil{
			fmt.Println("Something went wrong with the TUI")
			return
		}
		m := finalModel.(ui.AddModel)

		// Creating CSV file
		path, _ := filetool.CreateFile()

		// Open file / Close File
		file, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND, 0644)
		defer file.Close()
		if err != nil {
			fmt.Println("Something went wrong reading the file, please try again")
			return 
		}

		// Writing to file
		writer := csv.NewWriter(file)
		records := make([][]string, 0)
		count, err := filetool.GetLatestCount()
		records = append(records, []string{strconv.Itoa(count), m.Answers["Task Name"], m.Answers["Description"], "false"})
		err = writer.WriteAll(records)		

		fmt.Println("Task(s) added")
		return 

	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
