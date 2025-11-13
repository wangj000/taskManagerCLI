/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/wangj000/task/cmd"
	"fmt"
	ui "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	
	// TODO: i lowkey don't understand the whole casting thing
	
	// Initializes the new TUI cycle instance
	test := tea.NewProgram(ui.AddTUI())	

	// finalModel is the last instance of the model 
	// struct after cycle terminates
	finalModel, err := test.Run()

	if err != nil{
		fmt.Println("Something went wrong with the TUI")
		return
	}
	
	m := finalModel.(ui.Model)
	fmt.Println(m.Answers)

	cmd.Execute()
}
