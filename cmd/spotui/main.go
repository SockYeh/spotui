package main

import (
	"log"
	"spotui/internal/tui"
	"spotui/internal/tui/styles"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	
	if err := styles.LoadTheme("theme.toml"); err != nil {
		log.Fatal("Failed to load theme: ", err)
	}
	styles.Reload()

	if _, err := tea.NewProgram(tui.InitialModel()).Run(); err != nil {
		log.Fatal(err)
	}
}
