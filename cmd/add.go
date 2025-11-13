package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	filetool "github.com/wangj000/task/utils"
	"strconv"
	"os"
	"encoding/csv"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: `Append item
		- Single Item: task add [item]
		- Multiple Item: task add [item], [item]....`,
	Long: `Append item to current list`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Initializes the new TUI cycle instance
		addTUI := tea.NewProgram(ui.AddTUI())	

		// finalModel is the last instance of the model 
		// struct after cycle terminates
		finalModel, err := addTUI.Run()

		if err != nil{
			fmt.Println("Something went wrong with the TUI")
			return
		}
		
		m := finalModel.(ui.Model)

		// Creating the CSV file (if doesn't exist)
		path, _ := filetool.CreateFile()

		// Open file / Close File
		file, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND, 0644)
		defer file.Close()
		if err != nil {
			fmt.Println("Something went wrong reading the file, please try again")
			return 
		}

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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
