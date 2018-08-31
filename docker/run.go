package docker

import (
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

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
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: DohaImageLocal,
		Cmd: []string{"tail", "-f", "/dev/null"},
	}, nil, nil, "doha")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
}
