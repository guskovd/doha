package cmd

import (
	"github.com/spf13/cobra"
	"github.com/guskovd/doha/docker"
)

var start = &cobra.Command{
	Use:   "start [doha start]",
	Short: "Start doha daemon",
	Run: func(cmd *cobra.Command, args []string) {
		docker.BuildImage()
		docker.StartDaemonIfNotRunning()
	},
}

