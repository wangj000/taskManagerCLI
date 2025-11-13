package addTUI

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: Where I left off
// https://chatgpt.com/share/69141e1a-f714-8008-81f7-9be8d0eb3107

type model struct {
	ti textinput.Model
	questions []string
	q_index int 
	answers map[string]string
}

func AddTUI() model{
	
	ti := textinput.New()
	ti.Placeholder = "need to leetcode"
	ti.Focus()
	ti.CharLimit = 156;
	ti.Width = 20;

	return model {
		ti: ti,
		questions: []string{"Task Name:", "Description:"},
		q_index: 0,
		answers: make(map[string]string), 
	}

}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd){

	var cmd tea.Cmd

	switch msg := msg.(type){
		// If a key was pressed
	case tea.KeyMsg:
		switch msg.Type{
			// If the key entered is enter
		case tea.KeyEnter:
			if m.q_index < len(m.questions){
				question_answered := m.questions[m.q_index]
				m.answers[question_answered] = m.ti.Value()
			}
			m.q_index++
			// If the question index > len(questions) 
			// then the process is quit
			if m.q_index >= len(m.questions){
				return m, tea.Quit
			}
			
			// Deletes the previous view display
			m.ti.SetValue("")

		}

	}
	
	// Updates the previus view display with msg??
	// FIX: WTF IS HAPPENING HERE?
	m.ti, cmd = m.ti.Update(msg)
	return m, cmd 

}

func (m model) View() string {

	if m.q_index >= len(m.questions){
		return "DONE"
	}

	return fmt.Sprintf("%v\n%v\n\n(press Enter to submit)", 
		m.questions[m.q_index], m.ti.View())

}
