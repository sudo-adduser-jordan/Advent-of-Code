package styles

import (
	"encoding/json"
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func PrettyPrint(data interface{}) {
	var p []byte
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func Red(s string) lipgloss.Style {
	s = fmt.Sprintf(s)

	var style = lipgloss.NewStyle().
		SetString(s).
		Bold(true).
		Foreground(lipgloss.Color("6"))

	return style
}
