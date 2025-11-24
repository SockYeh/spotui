package main

import (
	"log"
	"spotui/internal/tui"
	"spotui/internal/tui/styles"
	"spotui/internal/utils"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	
	if err := styles.LoadTheme("theme.toml"); err != nil {
		log.Fatal("Failed to load theme: ", err)
	}
	styles.Reload()

	if err := utils.LoadConfig("config.toml"); err != nil {
		log.Fatal("Failed to load config: ", err)
	}

	if _, err := tea.NewProgram(tui.InitialModel(), tea.WithAltScreen()).Run(); err != nil {
		log.Fatal(err)
	}
}
