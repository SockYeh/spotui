package tui

import (
	"log"
	"spotui/internal/auth"
	"spotui/internal/tui/styles"
	"spotui/internal/utils"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct{
	width 		int
	height 		int
	currentView int
	accessToken string
}

func InitialModel() Model {

	if utils.Current.General.UseSpotify {
		resultChan, stop := auth.StartCallbackSever()

		result := <-resultChan
		if result.Error != nil {
			log.Fatalf("Error occured during spotify oauth grant: %s", result.Error)
		}
		stop()
		return Model{accessToken: result.AccessToken}
	}
	return Model{}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { 

	switch msg:=msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case tea.KeyRight.String(), "h":
			if m.currentView == 4 {
				m.currentView = 1
			} else {
				m.currentView++
			}
		case tea.KeyLeft.String(), "l":
			if m.currentView == 1 {
				m.currentView = 4
			} else {
				m.currentView--
			}
		}
	case tea.WindowSizeMsg:
		m.height, m.width = msg.Height, msg.Width
	}
	return m, nil 
}

func (m Model) View() string {
	var str strings.Builder
	styles.Reload()
	str.WriteString(renderMainView(m))
	return str.String()
}
