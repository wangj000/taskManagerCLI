package fileTools 

import (
	"fmt"
	"errors"
	"encoding/csv"
	"os"
	"io"
	"path/filepath"	
	"strconv"
	"slices"
)

// To create the CSV file
func CreateFile() (string, error) {
		
		path := filepath.Join("internal", "todos.csv")
		_, err := os.ReadFile(path)
		
		if err != nil{
			os.Create(path)	
		}

		return path, nil

}

// To get the last count item in the CSV 
func GetLatestCount() (int, error){

		path := filepath.Join("internal", "todos.csv")

		// Opening the file
		file, err := os.Open(path)
		if err != nil{
			return -1, errors.New("The file doesn't exist")
		}
		defer file.Close()
		
		// Reading the file line by line
		reader := csv.NewReader(file)
	
		prevCount := 0 

		for {

			record, err := reader.Read()

			if err == io.EOF {
				break
			}

			if err != nil {
				return -1, errors.New(fmt.Sprintf("%v error at reading file", err))
			}
		
			if _, isInt := strconv.Atoi(record[0]); isInt == nil{
				count, err := strconv.Atoi( record[0] )
				if err != nil {
					return -1, errors.New(fmt.Sprintf("%v error at int convert", err))
				}
				prevCount = count
			}

		}
		
		// Increment by one from the latest count
		prevCount += 1
		return prevCount, nil

}

// To filter tasks by taking in a slice of numbers and rebuilding 
// a array excluding the ones in the input parameter
func FilterTasks(ignoreTasks []string) ([][]string, error){
	
	path := filepath.Join("internal", "todos.csv")

	file, err := os.Open(path)
	if err != nil{
		return nil, errors.New("The file doesn't exist")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		return nil, errors.New("Something went wrong with reading the file")
	}

	remainingTasks := make([][]string, 0)
	
	for r := 0; r < len(records); r++ {

		if !slices.Contains(ignoreTasks, records[r][0]){
			remainingTasks = append(remainingTasks, records[r])
	 	}

	} 

	if len(remainingTasks) > 0{
		return remainingTasks, nil
	}

	return remainingTasks, nil 
}
