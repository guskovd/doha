package cmd

import (
	"github.com/spf13/cobra"
	"github.com/guskovd/doha/docker"
)

var shell = &cobra.Command{
	Use:   "shell [doha shell]",
	Short: "Invoke doha shell",
	Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		docker.BuildImage()
		docker.StartDaemonIfNotRunning()
		docker.ContainerExec()
	},
}

