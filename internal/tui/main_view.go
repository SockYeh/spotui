package tui

import (
	"spotui/internal/tui/styles"

	"github.com/charmbracelet/lipgloss"
)
func renderMainView(m Model) string {
	styles.CurrentView = m.currentView
	
	leftWidth := int(0.25*float64(m.width))
	centerWidth := int(0.5*float64(m.width))
	rightWidth := m.width - leftWidth - centerWidth - 6
	topHeight := int(float64(m.height)*0.7)
	headerHeight := int(float64(m.height)*0.05)

	header := styles.Title.Width(m.width - 2).Height(headerHeight).Render("SpoTUI - Playing through Spotify")
	left := styles.LeftPanel.Width(leftWidth).Height(topHeight).Render("Playlists")
	center := styles.CenterPanel.Width(centerWidth).Height(topHeight).Render("Playlist Info")
	right := styles.RightPanel.Width(rightWidth).Height(topHeight).Render("Queue")
	player := styles.PlayerBar.Width(m.width - 2).Height(m.height-topHeight-headerHeight-7).Render("Player")

	topRow := lipgloss.JoinHorizontal(lipgloss.Top, left, center, right)
	layout := lipgloss.JoinVertical(lipgloss.Left, header, topRow, player)

	return layout
}