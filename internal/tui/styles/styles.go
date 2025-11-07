package styles

import "github.com/charmbracelet/lipgloss"

var (
	LeftPanel   lipgloss.Style
	CenterPanel lipgloss.Style
	RightPanel  lipgloss.Style
	PlayerBar   lipgloss.Style
	Title       lipgloss.Style
)

func Reload() {
	var border lipgloss.Border
	switch Current.Styles.BorderStyle {
	case "rounded":
		border = lipgloss.RoundedBorder()
	case "double":
		border = lipgloss.DoubleBorder()
	default:
		border = lipgloss.NormalBorder()
	}

	base := lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Colors.Foreground)).
			Background(lipgloss.Color(Current.Colors.Background))

	PaneBase := base.Border(border).
			BorderForeground(lipgloss.Color(Current.Colors.Border)).
			Padding(Current.Styles.Padding).
			Margin(Current.Styles.Margin)

	LeftPanel = PaneBase
	CenterPanel = PaneBase.BorderForeground(lipgloss.Color(Current.Colors.Primary))
	RightPanel = PaneBase
	PlayerBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Current.Colors.Foreground)).
			Background(lipgloss.Color(Current.Colors.Background)).
			Padding(0,1)
	Title = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color(Current.Colors.Accent))
}
