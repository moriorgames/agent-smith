package smith

import (
	"testing"
	"github.com/docker/docker/client"
)

func TestContainerBuilderIsAbleToRemoveAndRunNewContainer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	container := new(Container)
	container.Name = "node_server"
	container.Image = "moriorgames/node-server"
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithVersion("v1.18"))
	BuildContainerByParams(container, cli)
}
