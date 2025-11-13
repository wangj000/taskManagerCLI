/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/wangj000/task/cmd"
	"fmt"
	TUI "github.com/wangj000/task/ui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	test := tea.NewProgram(TUI.AddTUI())	

	if _, err := test.Run(); err != nil{
		fmt.Println("Something went wrong with the TUI")
		return
	}

	cmd.Execute()
}
