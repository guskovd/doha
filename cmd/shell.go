package cmd

import (
	"github.com/spf13/cobra"
	// "os"
	// "log"
	// "io"
	// "io/ioutil"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	"github.com/guskovd/doha/docker"
)


func runDocker() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	imageName := "dguskov/doha:latest"

	cli.ImagePull(ctx, imageName, types.ImagePullOptions{})

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

}

var shell = &cobra.Command{
	Use:   "shell [doha shell]",
	Short: "Invoke doha shell",
	Long: `print is for printing anything back to the screen.
For many years people have printed back to the screen.`,
	Run: func(cmd *cobra.Command, args []string) {
		// runDocker()
		docker.BuildImage()
	},
}

