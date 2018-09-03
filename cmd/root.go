package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"fmt"
)

var rootCmd = &cobra.Command{
	Use:   "doha [doha cli]",
	Short: "DoHa cli",
	Long: `Docker-Habitat wrapper
                Complete documentation is available at https://github.com/guskovd/doha`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("hello from cobra-sample root")
	},
}


// Execute common
func Execute() {
	rootCmd.AddCommand(shell)
	rootCmd.AddCommand(exec)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

