package ui

import (
	"github.com/0xjuanma/golazo/internal/api"
	"github.com/0xjuanma/golazo/internal/constants"
	"github.com/0xjuanma/golazo/internal/data"
	"github.com/0xjuanma/golazo/internal/ui/design"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// BookmarksDialog is a dialog for selecting a club to bookmark.
type BookmarksDialog struct {
	match    api.Match
	selected int // 0 for home, 1 for away
}

// NewBookmarksDialog creates a new BookmarksDialog instance.
func NewBookmarksDialog(match api.Match) *BookmarksDialog {
	return &BookmarksDialog{
		match:    match,
		selected: 0,
	}
}

// ID returns the unique identifier of the dialog.
func (d *BookmarksDialog) ID() string {
	return "bookmarks"
}

// Update processes a message and returns the updated dialog and an optional action.
func (d *BookmarksDialog) Update(msg tea.Msg) (Dialog, DialogAction) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if d.selected > 0 {
				d.selected--
			}
		case "down", "j":
			if d.selected < 1 {
				d.selected++
			}
		case "enter":
			var team api.Team
			if d.selected == 0 {
				team = d.match.HomeTeam
			} else {
				team = d.match.AwayTeam
			}
			return d, BookmarksActionSelect{
				TeamID:   team.ID,
				TeamName: team.Name,
				LeagueID: d.match.League.ID,
			}
		case "esc":
			return d, DialogActionClose{}
		}
	}
	return d, nil
}

// BookmarksActionSelect is an action returned when a club is selected.
type BookmarksActionSelect struct {
	TeamID   int
	TeamName string
	LeagueID int
}

// View renders the dialog content.
func (d *BookmarksDialog) View(width, height int) string {
	contentWidth := 40
	contentHeight := 10

	settings, _ := data.LoadSettings()
	isHomeBookmarked := settings.IsClubBookmarked(d.match.HomeTeam.ID)
	isAwayBookmarked := settings.IsClubBookmarked(d.match.AwayTeam.ID)

	titleText := "Bookmark Club"
	promptText := " Select a club to bookmark:"
	if isHomeBookmarked || isAwayBookmarked {
		titleText = "Manage Bookmarks"
		promptText = " Manage your bookmarked clubs:"
	}

	title := design.RenderHeader(titleText, contentWidth)

	homeName := d.match.HomeTeam.Name
	awayName := d.match.AwayTeam.Name

	if isHomeBookmarked {
		homeName = "Remove ★ " + homeName
	} else {
		homeName = "Add ☆ " + homeName
	}
	if isAwayBookmarked {
		awayName = "Remove ★ " + awayName
	} else {
		awayName = "Add ☆ " + awayName
	}

	homeStyle := menuItemStyle
	awayStyle := menuItemStyle

	if d.selected == 0 {
		homeStyle = menuItemSelectedStyle
	} else {
		awayStyle = menuItemSelectedStyle
	}

	homeRow := "  " + homeStyle.Render(homeName)
	awayRow := "  " + awayStyle.Render(awayName)

	help := neonDimStyle.Width(contentWidth).Align(lipgloss.Center).Render(constants.HelpBookmarksDialog)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		"",
		promptText,
		"",
		homeRow,
		awayRow,
		"",
		help,
	)

	return neonPanelStyle.Width(contentWidth + 4).Height(contentHeight + 2).Render(content)
}
