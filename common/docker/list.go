package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func ListContainer(cli *client.Client) ([]types.Container, error) {
	ctx := context.Background()
	return cli.ContainerList(ctx, types.ContainerListOptions{})
}
