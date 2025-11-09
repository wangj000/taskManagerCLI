package createFile

import (
	"fmt"
	// "errors"
	"encoding/csv"
	"os"
	"path/filepath"	
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
