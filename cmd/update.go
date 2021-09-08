package cmd

import (
	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:     "update",
	Short:   "Updates a Passbolt Entity",
	Long:    `Updates a Passbolt Entity`,
	Aliases: []string{"change"},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}