package fileTools 

import (
	"fmt"
	"errors"
	"encoding/csv"
	"os"
	"io"
	"path/filepath"	
	"strconv"
)

func CreateFile() (string, error) {
		
		path := filepath.Join("internal", "todos.csv")
		_, err := os.ReadFile(path)
		
		if err != nil{

			os.Create(path)	

			file, _ := os.OpenFile(path, os.O_WRONLY | os.O_APPEND, 0644)

			writer := csv.NewWriter(file)
			
			new_headers := [][]string{{ "ID", "Task", "Status" }}

			err = writer.WriteAll(new_headers)

			if err != nil {
				fmt.Println("(createFile) -- Something went wrong writing to the file please try again.")
			}

			writer.Flush()

		}

		return path, nil

}

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

		return prevCount, nil

}

// TODO: THIS IS WHERE I LEFT OFF 
// func GetSiftCount() (int, error){
//
// }
