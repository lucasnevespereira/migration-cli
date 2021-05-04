package cmd

import (
	"fmt"
	"migration-cli/database"

	"github.com/spf13/cobra"
)

var migrateUpCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "up",
		Short: "migrate to v1 command",
		Long:  `Command to install version 1 of our application`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate up command")
			database.Open()
		},
	}

	migrateCmd.AddCommand(migrateUpCmd)
}
