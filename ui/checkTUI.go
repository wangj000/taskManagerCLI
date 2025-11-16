package ui 

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	CheckfocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153"))
	CheckcursorStyle = focusedStyle
	ChecktextStyle = focusedStyle
)

type CheckModel struct {
	ti textinput.Model	
	Answers string
}

func CheckTUI() CheckModel {

	ti := textinput.New()
	ti.TextStyle = DonefocusedStyle
	ti.Placeholder = "ex. 2"
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20;

	return CheckModel {
		ti: ti,
		Answers: "",
	}

}

func (m CheckModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CheckModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){

	var cmd tea.Cmd

	switch msg := msg.(type){
		case tea.KeyMsg:
			switch msg.Type {
				case tea.KeyEnter:
					m.Answers = m.ti.Value()
					return m, tea.Quit
			}
	}

	m.ti, cmd = m.ti.Update(msg)
	return m, cmd

}

func (m CheckModel) View() string {

		return fmt.Sprintf("Task ID\n%v\n\n (press Enter to checkoff", m.ti.View())

}
