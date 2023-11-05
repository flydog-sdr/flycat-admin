package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func RemoveContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{
		Force: true,
	})
}

func RemoveImage(cli *client.Client, id string) error {
	ctx := context.Background()
	_, err := cli.ImageRemove(ctx, id, types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	})

	return err
}
