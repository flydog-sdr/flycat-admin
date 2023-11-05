package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func StopContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerStop(ctx, id, container.StopOptions{})
}
