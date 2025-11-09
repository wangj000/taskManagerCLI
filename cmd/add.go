package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	creator "github.com/wangj000/task/internal"
	"strings"
	// "strconv"
	"os"
	"encoding/csv"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: `Append item
		- Single Item: task add [item]
		- Multiple Item: task add [item], [item]....`,
	Long: `Append item to current list`,
	Run: func(cmd *cobra.Command, args []string) {
		
		// Seperating multiple tasks by delimiter
		processed_string := strings.Join(args, "")
		processed_data := strings.Split(processed_string, ",")

		// Creating the CSV file (if doesn't exist)
		path, _ := creator.CreateFile()

		// Open file / Close File
		file, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND, 0644)

		defer file.Close()

		if err != nil {
			fmt.Println("Something went wrong reading the file, please try again")
			return 
		}

		// TODO: To track task number, you can read last latest in list, and then start relatively from that count.
		// - Edge case: if there are no prior item (i.e. You use multiple add first) then you should check for that.
		
		writer := csv.NewWriter(file)

		records := make([][]string, 0)

		for _, value := range processed_data{
			records = append(records, []string{"1", value, "completed!!!"})
		}

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
