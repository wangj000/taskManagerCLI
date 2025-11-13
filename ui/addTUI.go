package addTUI

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("153"))
	cursorStyle = focusedStyle
	textStyle = focusedStyle
)

type Model struct {
	ti textinput.Model
	questions []string
	q_index int 
	Answers map[string]string
}

func AddTUI() Model{
	
	// Initializes the internal UI buffer 
	ti := textinput.New()
	ti.TextStyle = focusedStyle
	ti.Placeholder = "Need to buy fruits" 
	ti.Focus()
	ti.CharLimit = 156;
	ti.Width = 20;

	return Model {
		ti: ti,
		questions: []string{"Task Name", "Description"},
		q_index: 0,
		Answers: make(map[string]string), 
	}

}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){

	var cmd tea.Cmd

	switch msg := msg.(type){
		// If a key was pressed
	case tea.KeyMsg:
		switch msg.Type{
			// If the key entered is enter
		case tea.KeyEnter:
			if m.q_index < len(m.questions){
				question_answered := m.questions[m.q_index]
				m.Answers[question_answered] = m.ti.Value()
			}
			m.q_index++
			// If the question index > len(questions) 
			// then the process is quit
			if m.q_index >= len(m.questions){
				return m, tea.Quit
			}

			// Deletes the internal text buffer by setting it to ""
			m.ti.SetValue("")

		}

	}
	
	// Updates the current interface depending on input
	// Note: msg is a rune slice that stores key inputs ( ex. []rune{'a', 'b', 'c'} ) 
	m.ti, cmd = m.ti.Update(msg)
	return m, cmd 

}

func (m Model) View() string {

	if m.q_index >= len(m.questions){
		return "DONE"
	}

	return fmt.Sprintf("%v\n%v\n\n(press Enter to submit)", 
		focusedStyle.Render(m.questions[m.q_index]), 
		m.ti.View())

}
