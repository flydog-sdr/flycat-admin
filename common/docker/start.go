package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func StartContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerStart(ctx, id, types.ContainerStartOptions{})
}
