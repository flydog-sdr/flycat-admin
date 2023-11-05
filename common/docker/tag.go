package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func TagImage(cli *client.Client, oldImage, newImage string) error {
	return cli.ImageTag(context.Background(), oldImage, newImage)
}
