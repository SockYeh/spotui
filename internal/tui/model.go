package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{
	currentView string
}

func InitialModel() Model { return Model{} }

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 
	switch msg:=msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil 
}

func (m Model) View() string {
	var str strings.Builder
	str.WriteString("SpoTUI - Press q to exit")
	str.WriteString("\n")

	switch m.currentView {
	case "main":
		str.WriteString(renderMainView(m))
	}
	
	return str.String()
}
