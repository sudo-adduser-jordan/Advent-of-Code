package styles

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// Labels
func PurpleLabel(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Background(lipgloss.Color("5"))
	return style
}
func BlueLabel(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Background(lipgloss.Color("6"))
	return style
}
func RedLabel(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Background(lipgloss.Color("1"))
	return style
}

// Text
func RedText(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(lipgloss.Color("1"))
	return style
}
func GreenText(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(lipgloss.Color("2"))
	return style
}
func BlueText(s string) lipgloss.Style {
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(lipgloss.Color("6"))
	return style
}

// Structs
func GreenStruct(t interface{}) lipgloss.Style {
	s := fmt.Sprintf("%+v", t)
	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(lipgloss.Color("2"))
	return style
}
