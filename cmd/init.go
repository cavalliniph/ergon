package cmd

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	projecttemplates "github.com/cavalliniph/arche/templates"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <project-directory>",
	Short: "Create a project from a template",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName, err := cmd.Flags().GetString("template")
		if err != nil {
			return err
		}

		if err := createProject(args[0], templateName); err != nil {
			return err
		}

		fmt.Fprintf(cmd.OutOrStdout(), "Project created at %s\n", args[0])
		return nil
	},
}

func init() {
	initCmd.Flags().StringP("template", "t", "flask", "project template")
	rootCmd.AddCommand(initCmd)
}

func createProject(destination, templateName string) error {
	if templateName != "flask" {
		return fmt.Errorf("unknown template %q (available: flask)", templateName)
	}

	if _, err := os.Stat(destination); err == nil {
		return fmt.Errorf("destination %q already exists", destination)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect destination %q: %w", destination, err)
	}

	templateRoot, err := fs.Sub(projecttemplates.Files, templateName)
	if err != nil {
		return fmt.Errorf("open template %q: %w", templateName, err)
	}

	if err := os.MkdirAll(destination, 0o755); err != nil {
		return fmt.Errorf("create destination %q: %w", destination, err)
	}

	copySucceeded := false
	defer func() {
		if !copySucceeded {
			_ = os.RemoveAll(destination)
		}
	}()

	err = fs.WalkDir(templateRoot, ".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		target := destination
		if path != "." {
			target = filepath.Join(destination, filepath.FromSlash(path))
		}

		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		data, err := fs.ReadFile(templateRoot, path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
	if err != nil {
		return fmt.Errorf("copy template %q: %w", templateName, err)
	}

	copySucceeded = true
	return nil
}
