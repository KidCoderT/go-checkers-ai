package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
}

func initialModel() model {
	return model{}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":

		// The "down" and "j" keys move the cursor down
		case "down", "j":

		// The "enter" key and the spacebar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":

		}

	case tea.MouseEvent:
		// if msg.Button == tea.MouseButton{
		// 	if msg.Action ==
		// }
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

var title = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#7D56F4")).
	Align(lipgloss.Center).
	Width(60).
	Height(3).
	AlignVertical(lipgloss.Center)

var RowStyle = lipgloss.NewStyle().
	Width(60).
	Align(lipgloss.Center)

var TileStyle = lipgloss.NewStyle().
	Bold(true).
	BorderStyle(lipgloss.RoundedBorder()).
	MarginRight(2).
	MarginBottom(1).
	PaddingLeft(1).
	PaddingRight(1).
	Width(4).
	AlignHorizontal(lipgloss.Center)

var whiteBlock = lipgloss.NewStyle().
	Inherit(TileStyle).
	Background(lipgloss.Color("#E8EDF9")).
	BorderForeground(lipgloss.Color("#E8EDF9")).
	BorderBackground(lipgloss.Color("#E8EDF9"))

var blackBlock = lipgloss.NewStyle().
	Inherit(TileStyle).
	BorderForeground(lipgloss.Color("63")).
	Background(lipgloss.Color("63")).
	BorderBackground(lipgloss.Color("63"))

func (m model) View() string {
	s := "\n"

	// Iterate over our choices
	// for i, choice := range m.choices {

	// 	// Is the cursor pointing at this choice?
	// 	cursor := " " // no cursor
	// 	if m.cursor == i {
	// 		cursor = ">" // cursor!
	// 	}

	// 	// Is this choice selected?
	// 	checked := " " // not selected
	// 	if _, ok := m.selected[i]; ok {
	// 		checked = "x" // selected!
	// 	}

	// 	// Render the row
	// 	s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	// }

	s += fmt.Sprintln(title.Render("Go Checkers!"))

	s += "\n\n"

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
	)))

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
	)))

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
	)))

	// 4th

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
	)))

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
		whiteBlock.Render(" "),
		blackBlock.Render(" "),
	)))

	// 6th

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
	)))

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
	)))

	s += fmt.Sprintln(RowStyle.Render(lipgloss.JoinHorizontal(
		lipgloss.Center,
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
		blackBlock.Render(" "),
		whiteBlock.Render("⚪"),
	)))

	s += "\n"

	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
