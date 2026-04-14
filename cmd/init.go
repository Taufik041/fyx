package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Taufik041/fyx/internal/config"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

// ── Palettes ────────────────────────────────────────────────────────────────

type palette struct {
	name     string
	selector string
	success  string
	dimmed   string
}

var palettes = []palette{
	{"Midnight", "#89DCEB", "#89B4FA", "#6C7086"},
	{"Amber", "#F9E2AF", "#89B4FA", "#6C7086"},
	{"Rose", "#F38BA8", "#89B4FA", "#6C7086"},
}

// ── Wizard state ─────────────────────────────────────────────────────────────

type step int

const (
	stepPalette step = iota
	stepProvider
	stepAPIKey
	stepHook
	stepDone
)

type wizardModel struct {
	step      step
	cursor    int
	pal       palette
	provider  string
	apiKey    string
	hook      bool
	textInput textinput.Model
	summary   []string
}

// ── Init the bubbletea model ─────────────────────────────────────────────────

func initialModel() wizardModel {
	ti := textinput.New()
	ti.EchoMode = textinput.EchoPassword
	ti.EchoCharacter = '•'
	ti.CharLimit = 200
	ti.Width = 50

	return wizardModel{
		step:      stepPalette,
		cursor:    0,
		pal:       palettes[0],
		textInput: ti,
	}
}

// ── Bubbletea interface — Init ───────────────────────────────────────────────

func (m wizardModel) Init() tea.Cmd {
	return nil
}

// ── Bubbletea interface — Update ─────────────────────────────────────────────

func (m wizardModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {

		case stepPalette:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < len(palettes)-1 {
					m.cursor++
				}
			case "enter":
				m.pal = palettes[m.cursor]
				m.summary = append(m.summary,
					fmt.Sprintf("  ✓ Theme      %s", m.pal.name),
				)
				m.step = stepProvider
				m.cursor = 0
			case "ctrl+c", "q":
				return m, tea.Quit
			}

		case stepProvider:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < 1 {
					m.cursor++
				}
			case "enter":
				if m.cursor == 0 {
					m.provider = "openai"
				} else {
					m.provider = "anthropic"
				}
				m.summary = append(m.summary,
					fmt.Sprintf("  ✓ Provider   %s", providerLabel(m.provider)),
				)
				m.step = stepAPIKey
				m.textInput.Placeholder = "Paste your API key here"
				m.textInput.Focus()
				return m, textinput.Blink
			case "ctrl+c", "q":
				return m, tea.Quit
			}

		case stepAPIKey:
			switch msg.String() {
			case "enter":
				key := strings.TrimSpace(m.textInput.Value())
				if key == "" {
					return m, nil
				}
				m.apiKey = key
				m.summary = append(m.summary,
					fmt.Sprintf("  ✓ API key    %s", maskKey(key)),
				)
				m.textInput.Blur()
				m.step = stepHook
				m.cursor = 0
			case "ctrl+c":
				return m, tea.Quit
			default:
				var cmd tea.Cmd
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			}

		case stepHook:
			switch msg.String() {
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < 1 {
					m.cursor++
				}
			case "enter":
				m.hook = m.cursor == 0
				hookLabel := "enabled"
				if !m.hook {
					hookLabel = "skipped"
				}
				m.summary = append(m.summary,
					fmt.Sprintf("  ✓ Hook       %s", hookLabel),
				)
				m.step = stepDone
				return m, tea.Quit
			case "ctrl+c", "q":
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

// ── Bubbletea interface — View ───────────────────────────────────────────────

func (m wizardModel) View() string {
	sel := lipgloss.NewStyle().Foreground(lipgloss.Color(m.pal.selector)).Bold(true)
	suc := lipgloss.NewStyle().Foreground(lipgloss.Color(m.pal.success))
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color(m.pal.dimmed))

	var b strings.Builder

	b.WriteString("\n")
	b.WriteString(sel.Render("  fyx — setup wizard") + "\n")
	b.WriteString(dim.Render("  ──────────────────────────────────────") + "\n")
	b.WriteString("\n")

	for _, line := range m.summary {
		b.WriteString(suc.Render(line) + "\n")
	}
	if len(m.summary) > 0 {
		b.WriteString("\n")
	}

	switch m.step {

	case stepPalette:
		b.WriteString("  Which theme would you like?\n\n")
		for i, p := range palettes {
			if i == m.cursor {
				b.WriteString(sel.Render("  ❯ "+p.name) + "\n")
			} else {
				b.WriteString(dim.Render("  • "+p.name) + "\n")
			}
		}
		b.WriteString("\n")
		b.WriteString(dim.Render("  ↑/↓ to move, enter to select") + "\n")

	case stepProvider:
		b.WriteString("  Which AI provider would you like to use?\n\n")
		providers := []string{"OpenAI", "Anthropic (Claude)"}
		for i, p := range providers {
			if i == m.cursor {
				b.WriteString(sel.Render("  ❯ "+p) + "\n")
			} else {
				b.WriteString(dim.Render("  • "+p) + "\n")
			}
		}
		b.WriteString("\n")
		b.WriteString(dim.Render("  ↑/↓ to move, enter to select") + "\n")

	case stepAPIKey:
		label := providerLabel(m.provider)
		b.WriteString(fmt.Sprintf("  Enter your %s API key:\n\n", label))
		b.WriteString("  " + m.textInput.View() + "\n")
		b.WriteString("\n")
		b.WriteString(dim.Render("  enter to confirm") + "\n")

	case stepHook:
		b.WriteString("\n  Install shell hook?\n")
		b.WriteString(dim.Render("  Enables automatic command correction when you mistype a command.") + "\n\n")
		options := []string{"Yes", "No"}
		for i, o := range options {
			if i == m.cursor {
				b.WriteString(sel.Render("  ❯ "+o) + "\n")
			} else {
				b.WriteString(dim.Render("  • "+o) + "\n")
			}
		}
		b.WriteString("\n")
		b.WriteString(dim.Render("  ↑/↓ to move, enter to select") + "\n")
	}

	return b.String()
}

// ── Helpers ──────────────────────────────────────────────────────────────────

func providerLabel(provider string) string {
	if provider == "anthropic" {
		return "Anthropic (Claude)"
	}
	return "OpenAI"
}

// Fix 1 — show first 2 chars + 6 dots + last 4 chars
func maskKey(key string) string {
	if len(key) <= 6 {
		return "••••••"
	}
	return key[:2] + "••••••" + key[len(key)-4:]
}

// ── Cobra command ─────────────────────────────────────────────────────────────

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "First time setup — configure your AI provider and API key",
	Run: func(cmd *cobra.Command, args []string) {
		runInit()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit() {
	m := initialModel()
	// tea.ClearScreen clears before the program starts
	p := tea.NewProgram(m, tea.WithAltScreen())

	result, err := p.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "  ✗ Wizard error: %v\n", err)
		os.Exit(1)
	}

	final, ok := result.(wizardModel)
	if !ok || final.step != stepDone {
		fmt.Println("\n  Setup cancelled.")
		return
	}

	// Fix 4 — clear screen before printing final summary so it shows only once
	// fmt.Print("\033[2J\033[H")

	// Save to config
	cfg := &config.Config{
		Provider: final.provider,
		APIKey:   final.apiKey,
		Active:   final.hook,
		Theme:    final.pal.name,
	}

	if err := config.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "  ✗ Failed to save config: %v\n", err)
		os.Exit(1)
	}

	sel := lipgloss.NewStyle().Foreground(lipgloss.Color(final.pal.selector)).Bold(true)
	suc := lipgloss.NewStyle().Foreground(lipgloss.Color(final.pal.success))
	dim := lipgloss.NewStyle().Foreground(lipgloss.Color(final.pal.dimmed))

	fmt.Println()
	fmt.Println(sel.Render("  fyx — setup wizard"))
	fmt.Println(dim.Render("  ──────────────────────────────────────"))
	fmt.Println()
	for _, line := range final.summary {
		fmt.Println(suc.Render(line))
	}
	fmt.Println()
	fmt.Println(sel.Render("  fyx is ready. Try: fyx browse kubectl"))
	fmt.Println(dim.Render("  ──────────────────────────────────────"))
	fmt.Println()
}
