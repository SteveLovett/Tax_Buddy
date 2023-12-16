package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	prompt   string
	choices  []string
	cursor   int
	selected map[int]struct{}
}

type personOnReturn interface {
	personalInformation
}

type personalInformation struct {
	FirstName string
	LastName  string
	SSN       string
	DOB       string
	Email     string
	Phone     string
	Address   string
	City      string
	State     string
	Zip       string
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error: ", err)
		os.Exit(1)
	}
}

func initialModel() model {
	return model{
		choices:  []string{"US Address", "Foreign Address"},
		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return boxset(m, msg)
}

func (m model) View() string {
	m.prompt = "Please enter the taxpayer's personal information:\n\n"

	mPointer := &m
	boxsetView(mPointer)

	m.prompt += "\nPress q to quit.\n"
	return m.prompt
}
