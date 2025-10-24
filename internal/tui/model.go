package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct{}

func InitialModel() model { return model{} }

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 
	switch msg:=msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil 
}

func (m model) View() string {
	var str strings.Builder
	str.WriteString("SpoTUI - Press q to exit")
	str.WriteString("\n")
	return str.String()
}
