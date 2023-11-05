package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func DeployContainer(cli *client.Client, options *DeployOptions) error {
	ctx := context.Background()

	resp, err := cli.ContainerCreate(
		ctx, options.Container, options.Host,
		options.Network, nil, options.Name,
	)
	if err != nil {
		return err
	}

	return cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
}
