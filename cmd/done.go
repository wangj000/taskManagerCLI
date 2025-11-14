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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string){

		// Initializes the new TUI cycle instance
		doneTUI := tea.NewProgram(ui.DoneTUI())	

		// finalModel is the last instance of the model 
		// struct after cycle terminates
		finalModel, err := doneTUI.Run()
		if err != nil{
			fmt.Println("Something went wrong with the TUI")
			return
		}
		
		m := finalModel.(ui.DoneModel)
		path := filepath.Join("internal", "todos.csv")

		del_data := strings.Split(m.Answers, "")

		fmt.Println(del_data)
		
		cur_data, err := filetool.FilterTasks(del_data)
		if err != nil{
			fmt.Println("Something went wrong filtering the tasks.")	
			fmt.Println(err)
			return
		}

		count	:= 0	
		for i := 0 ; i < len(cur_data); i++ {
			count += 1
			cur_data[i][0] = strconv.Itoa(count)
		}

		err = os.Remove(path)	
		if err != nil {
			fmt.Println("Something went wrong removing the file when trying to rewrite")
			return
		}

		_, err = filetool.CreateFile()
		if err != nil {
			fmt.Println("Something went wrong trying to recreate the file when trying to rewrite")
		}
		
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

		writer := csv.NewWriter(file)

		err = writer.WriteAll(cur_data)		
		if err != nil{
			fmt.Println("Something went wrong writing to the file")
		}
		
		// Close file
		return

	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
