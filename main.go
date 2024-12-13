package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type song struct {
	title  string
	artist string
	album  string
}

type model struct {
	songs []song
}

// I could as well just init it somewhere else (maybe)
func initModel() model {
	return model{
		songs: []song{
			{
				title:  "Song for the Dead",
				artist: "Queens of the Stone Age",
				album:  "Queens of the Stone Age",
			},
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Background(lipgloss.Color("#7E9CD8")).
		Foreground(lipgloss.Color("#54546D")).
		Padding(0, 2)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#DCD7BA"))

	s := lipgloss.PlaceHorizontal(0, lipgloss.Center, titleStyle.Render("MTUI"))

	for _, song := range m.songs {
		s += textStyle.Render(fmt.Sprintf("\n\n%s\n[%s]\n%s\n", song.title, song.artist, song.album))
	}
	return s
}

func main() {
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there has been an error: %v", err)
		os.Exit(1)
	}
}
