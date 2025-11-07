package styles

import "github.com/charmbracelet/lipgloss"

var (
	PrimaryText lipgloss.Style
	PanelBox    lipgloss.Style
)

func Reload() {
	PrimaryText = lipgloss.NewStyle().
		Foreground(lipgloss.Color(Current.Colors.Primary))

	PanelBox = lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(Current.Colors.Border)).
		Padding(Current.Styles.Padding).
		Margin(Current.Styles.Margin)
}
