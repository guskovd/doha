package docker

import (
	"fmt"
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

	_, err1 := cli.ContainerInspect(context.Background(), "doha")
	if err1 != nil {
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

	resp, err := cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: DohaImageLocal,
			Cmd: []string{"tail", "-f", "/dev/null"},
			User: fmt.Sprintf("%s:%s", current_user.Uid, current_user.Gid),
			// Volumes: map[string]struct{}{"/hab:/hab": {}},
		},
		&container.HostConfig{
			Mounts: []mount.Mount{
				{
					Type:   mount.TypeBind,
					Source: "/hab",
					Target: "/hab",
				},
				{
					Type:   mount.TypeBind,
					Source: os.Getenv("HOME"),
					Target: os.Getenv("HOME"),
				},
			},
		},
		nil,
		"doha",
	)
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
