package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{
	width int
	height int
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
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	}
	return m, nil 
}

func (m Model) View() string {
	var str strings.Builder
	str.WriteString(renderMainView(m))
	return str.String()
}
