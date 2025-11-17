package cmd

import (
	"fmt"
	filetool "github.com/wangj000/task/utils"
	"github.com/spf13/cobra"
	"path/filepath"	
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "To delete a item from the list",
	Long: `To delete a item from the table by using 'task done' 
	and typing in the task ID upon being prompting by the TUI`,
	Run: func(cmd *cobra.Command, args []string){

		// Initializes the new TUI program instance 
		doneTUI := tea.NewProgram(ui.DoneTUI())	

		// Intializes the TUI program
		finalModel, err := doneTUI.Run()
		if err != nil{
			fmt.Println("Something went wrong with the TUI")
			return
		}
		
		// The latest struct post program termination
		m := finalModel.(ui.DoneModel)
		path := filepath.Join("internal", "todos.csv")

		// Filters the deleted item   
		del_data := strings.Split(m.Answers, "")
		cur_data, err := filetool.FilterTasks(del_data)
		if err != nil{
			fmt.Println("Something went wrong filtering the tasks.")	
			fmt.Println(err)
			return
		}

		// Recount the items in the list ID
		count	:= 0	
		for i := 0 ; i < len(cur_data); i++ {
			count += 1
			cur_data[i][0] = strconv.Itoa(count)
		}

		// Delete the CSV
		err = os.Remove(path)	
		if err != nil {
			fmt.Println("Something went wrong removing the file when trying to rewrite")
			return
		}

		// Recreating the file
		_, err = filetool.CreateFile()
		if err != nil {
			fmt.Println("Something went wrong trying to recreate the file when trying to rewrite")
		}
		
		// Opening the file for writes
		file, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Something went wrong opening the file")
			return
		}

		if _, err := os.Stat(path); err != nil {
			fmt.Println("%v", err)
			return
		}
		defer file.Close()

		// Writing tot he file
		writer := csv.NewWriter(file)
		err = writer.WriteAll(cur_data)		
		if err != nil{
			fmt.Println("Something went wrong writing to the file")
			return
		}
		
		// Close file
		return

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
