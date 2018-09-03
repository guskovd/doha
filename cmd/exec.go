package cmd

import (
	"github.com/spf13/cobra"
	"github.com/guskovd/doha/docker"
)

var exec = &cobra.Command{
	Use:   "exec [doha exec]",
	Short: "Invoke doha exec commmand",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		docker.BuildImage()
		docker.StartDaemonIfNotRunning()
		docker.ContainerExec(args)
	},
}

