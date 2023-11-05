package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func RestartContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerRestart(ctx, id, container.StopOptions{})
}
