package docker

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

	mount "github.com/docker/docker/api/types/mount"
)

func StartDaemonIfNotRunning() {
	cli, err := client.NewEnvClient()

	if err != nil {
		panic(err)
	}

	resp, err := cli.ContainerInspect(context.Background(), DohaContainerName)

	if err != nil {
		RunContainer()
	} else if resp.State.Running == false {
		log.Println("Doha daemon in Dead state. Restarting ...")
		StopContainer()
		RunContainer()
	}
}

func RunContainer() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	current_user, err := user.Current()

	if err != nil {
		panic(err)
	}

	log.Println("Starting doha daemon ...")
	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: DohaImageLocal,
			User:  fmt.Sprintf("%s:%s", current_user.Uid, current_user.Gid),
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeVolume,
					Source: "hab",
					Target: "/hab",
				},
				{
					Type:   mount.TypeBind,
					Source: "/var/run/docker.sock",
					Target: "/var/run/docker.sock",
				},
				{
					Type:   mount.TypeBind,
					Source: os.Getenv("HOME"),
					Target: os.Getenv("HOME"),
				},
			},
		},
		nil,
		DohaContainerName,
	)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
