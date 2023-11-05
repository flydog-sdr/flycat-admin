package docker

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/jsonmessage"
)

func PullImage(cli *client.Client, name string) error {
	ctx := context.Background()

	responseBody, err := cli.ImagePull(
		ctx, name, types.ImagePullOptions{},
	)
	if err != nil {
		return err
	}
	defer responseBody.Close()

	decoder := json.NewDecoder(responseBody)
	for {
		var message jsonmessage.JSONMessage
		if err := decoder.Decode(&message); err != nil {
			if err == io.EOF {
				break
			}

			return err
		}

		if message.Progress != nil && message.Progress.String() != "" {
			log.Printf("%s: %s", message.ID, message.ProgressMessage)
		} else if message.Status != "" {
			log.Println(strings.TrimSpace(message.Status))
		}
	}

	return nil
}
