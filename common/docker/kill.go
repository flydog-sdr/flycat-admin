package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func KillContainer(cli *client.Client, id string) error {
	ctx := context.Background()
	return cli.ContainerKill(ctx, id, "SIGKILL")
}
