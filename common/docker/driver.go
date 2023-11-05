package docker

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/docker/docker/client"
)

func DockerConnect(socket string) (*client.Client, error) {
	cli, err := client.NewClientWithOpts([]client.Opt{
		client.FromEnv, client.WithHTTPClient(&http.Client{
			Transport: &http.Transport{
				DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
					return net.Dial("unix", socket)
				},
			},
		}), client.WithHost(fmt.Sprintf("unix://%s", socket)),
	}...)
	if err != nil {
		return nil, err
	}

	return cli, nil
}

func DockerClose(c *client.Client) error {
	return c.Close()
}
