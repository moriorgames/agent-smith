package smith

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

//"github.com/docker/go-connections/nat"

func BuildContainerByParams(deploy *Container, cli *client.Client) {
	ctx := context.Background()
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	removeExistingContainer(deploy, containers, cli)

	out, err := cli.ImagePull(ctx, deploy.Image, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, out)

	//ports := []string{"8000/tcp"}
	//exposedPorts, _, err := nat.ParsePortSpecs(ports)
	//containerConfig := &container.Config{
	//	Image:        deploy.Image,
	//	ExposedPorts: exposedPorts,
	//}
	containerConfig := &container.Config{
		Image: deploy.Image,
	}
	resp, err := cli.ContainerCreate(ctx, containerConfig, nil, nil, deploy.Name)
	if err != nil {
		panic(err)
	}
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.ID)
}

func removeExistingContainer(container *Container, containers []types.Container, cli *client.Client) {
	for _, data := range containers {
		for _, name := range data.Names {
			containerName := strings.Replace(name, "/", "", -1)
			fmt.Println(containerName, data.ID)
			if containerName == container.Name {
				err := cli.ContainerRemove(context.Background(), data.ID, types.ContainerRemoveOptions{
					Force: true,
				})
				fmt.Println(err)
			}
		}
	}
}
