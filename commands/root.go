/*
Copyright © 2023 libraria-app
*/
package commands

import (
	"github.com/libraria-app/cli/internal/librariacli"
	"github.com/spf13/cobra"
)

func NewRootCommand(lcli *librariacli.Lcli) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "libraria",
		Short: "The Libraria CLI client",
		Long: `This app is a Libraria CLI client that help you to manage your translations just from your terminal.
It allows developer to export term translation from the needed project.`,
	}

	rootCmd.AddCommand(newExportCommand(lcli))

	return rootCmd
}
