/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"charm.land/huh/v2"
)

var (
	burger string
	confirm bool
)

// flaskCmd represents the flask command
var flaskCmd = &cobra.Command{
	Use:   "flask",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Escolha o projeto").
					Options(
						huh.NewOption("Backend (Python + Flask)", "back"),
						huh.NewOption("Frontend (React + Vite)", "front"),
						huh.NewOption("Mobile (React-Native)", "mobile"),
					).
					Value(&burger),
			),
		).WithTheme(huh.ThemeFunc(huh.ThemeBase))

		err := form.Run()

		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	createCmd.AddCommand(flaskCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// flaskCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// flaskCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	flaskCmd.Flags().String("name", "", "Cria o projeto com o nome que desejar")
}
