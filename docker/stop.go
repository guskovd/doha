package docker

import (
	"log"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"

)

func StopContainer() {
	log.Println("Terminating doha daemon ...")

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStop(ctx, DohaContainerName, nil); err != nil {
		log.Println(fmt.Sprintf("Container stop: %s", err))
	}

	if err := cli.ContainerRemove(ctx, DohaContainerName, types.ContainerRemoveOptions{}); err != nil {
		log.Println(fmt.Sprintf("Container delete: %s", err))
	}
}

