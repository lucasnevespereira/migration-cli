package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var migrateDownCmd *cobra.Command

func init() {
	migrateUpCmd = &cobra.Command{
		Use:   "down",
		Short: "migrate from v2 to v1",
		Long:  `Command to downgrade database from v2 to v1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Running migrate down command")
		},
	}

	migrateCmd.AddCommand(migrateDownCmd)
}
