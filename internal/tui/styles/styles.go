package styles

import "github.com/charmbracelet/lipgloss"

var (
	LeftPanel   lipgloss.Style
	CenterPanel lipgloss.Style
	RightPanel  lipgloss.Style
	PlayerBar   lipgloss.Style
	Title       lipgloss.Style
	CurrentView int
)

var viewTable map[int]*lipgloss.Style

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
			Foreground(lipgloss.Color(Current.Colors.Foreground))

	PaneBase := base.Border(border).
			BorderForeground(lipgloss.Color(Current.Colors.Border)).
			Padding(Current.Styles.Padding)
	

	LeftPanel = PaneBase
	CenterPanel = PaneBase
	RightPanel = PaneBase
	PlayerBar = PaneBase.
			Foreground(lipgloss.Color(Current.Colors.Foreground)).
			Padding(0,1)

	viewTable = map[int]*lipgloss.Style{
		1: &LeftPanel,
		2: &CenterPanel,
		3: &RightPanel,
		4: &PlayerBar,
	}

	if s, ok := viewTable[CurrentView]; ok{
		
		*s = s.BorderForeground(lipgloss.Color(Current.Colors.Primary))
	}

	Title = PaneBase.
			Bold(true).
			Foreground(lipgloss.Color(Current.Colors.Accent)).
			Background(lipgloss.Color(Current.Colors.Background)).
			Align(lipgloss.Center).
			Padding(1, 2)

}
