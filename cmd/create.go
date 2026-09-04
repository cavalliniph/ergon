/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	choices []string
	cursor int
	selected string
}

func initialModel() model {
	return model{
		choices: []string{
			"[Backend] Python + Flask",
			"[Frontend] React + Vite",
			"[Mobile] React Native",
		},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices) -1 {
				m.cursor++
			}

		case "enter", "space":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() tea.View {
	s := "Escolha seu tipo de projeto\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if m.selected == choice {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nPressione q para sair.\n"

	return tea.NewView(s)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p := tea.NewProgram(initialModel())
		finalModel, err := p.Run()

		if err != nil {
			fmt.Printf("um erro ocorreu: %v", err)
			os.Exit(1)
		}

		result := finalModel.(model)
		
		fmt.Printf("chosen > %v\n", result.selected)
		
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
