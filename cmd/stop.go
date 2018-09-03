package cmd

import (
	"github.com/spf13/cobra"
	"github.com/guskovd/doha/docker"
)

var stop = &cobra.Command{
	Use:   "stop [doha stop]",
	Short: "Stop doha daemon",
	Run: func(cmd *cobra.Command, args []string) {
		docker.StopContainer()
	},
}

