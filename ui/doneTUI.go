package ui 

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
	tea "github.com/charmbracelet/bubbletea"
)

var (
	DonefocusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153"))
	DonecursorStyle = focusedStyle
	DonetextStyle = focusedStyle
)

type DoneModel struct {
	ti textinput.Model
	Answers string 
}

func DoneTUI() DoneModel {

	ti := textinput.New()
	ti.TextStyle = DonefocusedStyle
	ti.Placeholder = "ex. 2"
	ti.Focus()
	ti.CharLimit = 3
	ti.Width = 20;

	return DoneModel {
		ti: ti,
		Answers: "",
	}

}

	func (m DoneModel) Init() tea.Cmd {
		return textinput.Blink
	}

	func (m DoneModel) Update(msg tea.Msg) (tea.Model, tea.Cmd){

		var cmd tea.Cmd
	
		switch msg := msg.(type){

			case tea.KeyMsg:

				switch msg.Type {
					case tea.KeyEnter:
						// Gets the text buffer value for the 
						// current text that the user typed out
						m.Answers = m.ti.Value()	
						return m, tea.Quit
				}

		}
	
		m.ti, cmd = m.ti.Update(msg)	
		return m, cmd

	}

	func (m DoneModel) View() string {
		
		return fmt.Sprintf("%v\n%v\n\n(press Enter to delete", DonefocusedStyle.Render("TASK ID"), m.ti.View())  

	}

