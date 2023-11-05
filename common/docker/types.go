package docker

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

type DeployOptions struct {
	Name      string
	Container *container.Config
	Network   *network.NetworkingConfig
	Host      *container.HostConfig
}
