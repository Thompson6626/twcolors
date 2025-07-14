package main

import (
	"context"
	"fmt"
	"image/color"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
	"github.com/spf13/cobra"
	"golang.design/x/clipboard"
)

var (
	Version = "1"

	rootCmd = &cobra.Command{
		Use:     "twcolors",
		Short:   "Explore and copy Tailwind CSS colors in your terminal",
		Long:    `A TUI for browsing the Tailwind CSS color palette and copying HEX or OKLCH values.`,
		Version: Version,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := clipboard.Init(); err != nil {
				return fmt.Errorf("failed to initialize clipboard: %w", err)
			}

			p := tea.NewProgram(initialModel(), tea.WithAltScreen())
			if _, err := p.Run(); err != nil {
				return fmt.Errorf("TUI error: %w", err)
			}
			return nil
		},
	}
)

var familyOrder = []string{
	"RED", "ORANGE", "AMBER", "YELLOW", "LIME", "GREEN", "EMERALD",
	"TEAL", "CYAN", "SKY", "BLUE", "INDIGO", "VIOLET", "PURPLE",
	"FUCHSIA", "PINK", "ROSE", "SLATE", "GRAY", "ZINC", "NEUTRAL", "STONE",
}

var shadeLevels = []int{50, 100, 200, 300, 400, 500, 600, 700, 800, 900, 950}

var blends = gamut.Blends(lipgloss.Color("#00C9FF"), lipgloss.Color("#92FE9D"), 50)

type Color struct {
	Family string
	Shade  int
	Hex    lipgloss.Color
}

type Model struct {
	width, height    int
	cursorX, cursorY int
	clipboard        string
	clipboardTime    time.Time
	buffer           string
	boxStyle         lipgloss.Style
	hideLabels       bool
	currentFormat    ColorType
}

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt, syscall.SIGTERM,
	)
	defer cancel()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initialModel() Model {
	return Model{
		boxStyle: lipgloss.NewStyle().
			Width(6).
			Height(2).
			Align(lipgloss.Center, lipgloss.Center).
			Padding(0, 2),
		currentFormat: HEX,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m.handleResize(msg)
	case tea.KeyMsg:
		return m.handleKey(msg)
	}
	return m, nil
}

func (m *Model) handleResize(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height

	boxW := 6
	boxH := 2
	pad := 2

	if msg.Width < 138 {
		boxW = 5
		pad = 1
	}

	if msg.Width < 120 {
		boxW = 4
	}

	if msg.Width < 50 {
		boxW = 2
	}

	if msg.Width < 30 {
		boxW = 1
	}

	if msg.Height < 26 {
		boxH = 1
	}

	m.boxStyle = lipgloss.NewStyle().
		Width(boxW).
		Height(boxH).
		Align(lipgloss.Center, lipgloss.Center).
		Padding(0, pad)

	return m, nil
}

func (m *Model) handleKey(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "esc":
		return m, tea.Quit
	case "left":
		if m.cursorX > 0 {
			m.cursorX--
		}
	case "right":
		if m.cursorX < len(familyOrder)-1 {
			m.cursorX++
		}
	case "up":
		if m.cursorY > 0 {
			m.cursorY--
		}
	case "down":
		if m.cursorY < len(shadeLevels)-1 {
			m.cursorY++
		}
	case "tab":
		if m.currentFormat == HEX {
			m.currentFormat = OKLCH
		} else {
			m.currentFormat = HEX
		}
	case "a", "b", "c", "d", "e", "f", "g", "h", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v":
		m.buffer = msg.String()
	case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ":":
		if m.buffer != "" {
			move := m.buffer + msg.String()
			m.buffer = ""
			return m.jumpToCursor(move)
		}
	case "enter":
		color := m.getColor(m.cursorX, m.cursorY)
		colorValue := Palette[color.Family][color.Shade][m.currentFormat]
		clipboard.Write(clipboard.FmtText, []byte(colorValue))
		m.clipboard = fmt.Sprintf("%s", colorValue)
		m.clipboardTime = time.Now()
	case " ": // Spacebar
		m.hideLabels = !m.hideLabels
	}
	return m, nil
}

func (m *Model) jumpToCursor(pos string) (tea.Model, tea.Cmd) {
	if len(pos) != 2 {
		return m, nil
	}
	colRune := rune(pos[0])
	if colRune < 'a' || colRune > 'v' {
		return m, nil
	}
	col := int(colRune - 'a')

	var row int
	if pos[1] == ':' {
		row = 10
	} else if pos[1] >= '0' && pos[1] <= '9' {
		row = int(pos[1] - '0')
	} else {
		return m, nil
	}

	if col >= len(familyOrder) || row >= len(shadeLevels) {
		return m, nil
	}

	m.cursorX = col
	m.cursorY = row
	return m, nil
}

func (m Model) getColor(x, y int) Color {
	return Color{
		Family: familyOrder[x],
		Shade:  shadeLevels[y],
		Hex:    lipgloss.Color(Palette[familyOrder[x]][shadeLevels[y]][HEX]),
	}
}

func (m Model) View() string {
	var b strings.Builder

	for y, shade := range shadeLevels {
		var row []string
		for x, family := range familyOrder {
			color := Color{
				Family: family,
				Shade:  shade,
				Hex:    lipgloss.Color(Palette[family][shade][HEX]),
			}
			selected := m.cursorX == x && m.cursorY == y
			row = append(row, renderBox(color, selected, y, x, m.boxStyle, m.hideLabels))
		}
		b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, row...) + "\n")
	}

	info := ""
	if m.clipboard != "" && time.Since(m.clipboardTime) < 3*time.Second {
		info = lipgloss.NewStyle().
			MarginTop(1).
			Render(rainbow(lipgloss.NewStyle(), fmt.Sprintf("%s copied to clipboard", m.clipboard), blends))
	}

	formatLabelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))       // white
	formatValueStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#34D399")) // emerald

	formatStatus := lipgloss.NewStyle().MarginTop(1).Render(
		formatLabelStyle.Render("Clipboard format: ") + formatValueStyle.Render(formatName(m.currentFormat)),
	)

	// Suboptimal
	gray := lipgloss.NewStyle().Foreground(lipgloss.Color("7"))
	keyStyle := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFA500"))
	hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	hintText := func(s string) string { return hintStyle.Render(s) }
	key := func(s string) string { return keyStyle.Render(s) }

	line1 := gray.Render(
		hintText("Use ") + key("← ↑ → ↓") + hintText(" to move   ") +
			hintText("Type ") + key("coordinates") + hintText(" (e.g. a0)   ") +
			hintText("Press ") + key("Enter") + hintText(" to copy"),
	)

	line2 := gray.Render(
		hintText("Press ") + key("Spacebar") + hintText(" to toggle labels   ") +
			hintText("Press ") + key("Tab") + hintText(" to toggle ") + key("HEX/OKLCH") + "   " +
			hintText("Press ") + key("Esc") + hintText(" to quit"),
	)

	hint := lipgloss.NewStyle().
		MarginTop(1).
		Align(lipgloss.Center).
		Render(line1 + "\n" + line2)

	return lipgloss.Place(
		m.width, m.height,
		lipgloss.Center, lipgloss.Center,
		b.String()+info+"\n"+formatStatus+"\n"+hint,
	)
}

func renderBox(c Color, selected bool, y, x int, boxStyle lipgloss.Style, hideLabels bool) string {
	fg := lipgloss.Color("#000")
	if c.Shade > 500 {
		fg = "#fff"
	}

	label := ""
	if !hideLabels {
		label = getLabel(x, y)
	}
	if selected {
		label = "⬤"
	}

	return boxStyle.Background(c.Hex).Foreground(fg).Render(label)
}

func getLabel(x, y int) string {
	letter := string(rune('a' + x))
	index := ":"
	if y < 10 {
		index = strconv.Itoa(y)
	}
	return letter + index
}

func rainbow(base lipgloss.Style, s string, colors []color.Color) string {
	var str string
	for i, ss := range s {
		colorr, _ := colorful.MakeColor(colors[i%len(colors)])
		str += base.Foreground(lipgloss.Color(colorr.Hex())).Render(string(ss))
	}
	return str
}

func formatName(ct ColorType) string {
	switch ct {
	case HEX:
		return "HEX"
	case OKLCH:
		return "OKLCH"
	default:
		return "UNKNOWN"
	}
}
