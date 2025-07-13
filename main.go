package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	boxW = 4
	boxH = 2
)

var (
	boxStyle = lipgloss.NewStyle().
		Width(boxW).
		Height(boxH).
		Align(lipgloss.Center, lipgloss.Center).Padding(0, 0)

	familyOrder = []string{
		"RED", "ORANGE", "AMBER", "YELLOW", "LIME", "GREEN", "EMERALD",
		"TEAL", "CYAN", "SKY", "BLUE", "INDIGO", "VIOLET", "PURPLE",
		"FUCHSIA", "PINK", "ROSE", "SLATE", "GRAY", "ZINC", "NEUTRAL", "STONE",
	}

	shadeLevels = []int{950, 900, 800, 700, 600, 500, 400, 300, 200, 100, 50}
)

type Model struct {
	width  int
	height int
	buffer string
}

func main() {
	m := &Model{}
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.MouseMsg:
		if msg.Button != tea.MouseButtonLeft {
			return m, nil
		}

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v":
			m.buffer = msg.String()
		case "1", "2", "3", "4", "5", "6", "7", "8", "9":

		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}
	return m, nil
}

func (m *Model) View() string {
	var b strings.Builder

	// Sort families for consistency
	var families []string
	for fam := range Palette {
		families = append(families, fam)
	}

	// For each shade level, create a row
	for _, shade := range shadeLevels {
		var row []string
		for _, family := range familyOrder {
			color := Palette[family][shade]
			fg := lipgloss.Color("#000") // black
			if shade > 500 {
				fg = "#fff" // white for dark backgrounds
			}
			cell := boxStyle.Background(color).Foreground(fg).Render("")
			row = append(row, cell)
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...) + "\n")
	}

	grid := b.String()

	// Center the grid
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		grid,
	)
}
