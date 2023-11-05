package docker

import (
	"context"

	"github.com/docker/docker/client"
)

func GetContainerID(cli *client.Client, name string) (string, error) {
	container, err := cli.ContainerInspect(context.Background(), name)
	return container.ID, err
}

func GetImageID(cli *client.Client, name string) (string, error) {
	image, _, err := cli.ImageInspectWithRaw(context.Background(), name)
	return image.ID, err
}
