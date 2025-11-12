package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"io"
	"encoding/csv"
	"path/filepath"	
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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
		
		// Reading the file line by line
		reader := csv.NewReader(file)
		rows := make([][]string, 0)
		for {

			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("%v second print - Something went wrong reading the file, please try again", err)
				return
			}
			
			row_items := make([]string, 0)
			for _, item := range record {
				row_items = append(row_items, item)
			}

			rows = append(rows, row_items)

		}
			
		var (
			purple    = lipgloss.Color("153")
			gray      = lipgloss.Color("245")
			lightGray = lipgloss.Color("153")

			headerStyle  = lipgloss.NewStyle().
				Foreground(purple).
				Bold(true).
				Align(lipgloss.Center)

			cellStyle    = lipgloss.NewStyle().
				Padding(0, 1).
				Width(30).
				BorderTop(false).      
				BorderBottom(false).   
				BorderLeft(true).      
				BorderRight(true)
			oddRowStyle  = cellStyle.Foreground(gray)
			evenRowStyle = cellStyle.Foreground(lightGray)
		)

		t := table.New().
				Border(lipgloss.NormalBorder()).
				BorderStyle(lipgloss.NewStyle().Foreground(lightGray)).
				StyleFunc(func(row, col int) lipgloss.Style {
						switch {
						case row == table.HeaderRow:
								return headerStyle
						case row%2 == 0:
								return evenRowStyle
						default:
								return oddRowStyle
						}
				}).
				Headers("ID", "Name", "Description","Completed").
				Rows(rows...)

		// You can also add tables row-by-row
		
		fmt.Println(t)

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
