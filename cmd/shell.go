package cmd

import (
	"github.com/spf13/cobra"
	"github.com/guskovd/doha/docker"
)

var shell = &cobra.Command{
	Use:   "shell [doha shell]",
	Short: "Invoke doha shell",
	Run: func(cmd *cobra.Command, args []string) {
		docker.BuildImage()
		docker.StartDaemonIfNotRunning()
		docker.ContainerExec([]string{"/bin/bash"})
	},
}

