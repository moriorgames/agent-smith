package smith

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
)

type TableRow struct {
	ContainerImage     string
	ContainerCreatedAt string
	ContainerPorts     string
}

func ViewHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, renderHomeContent())
}

func renderHomeContent() string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)

	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	var rows []TableRow

	for _, container := range containers {

		createdAt := time.Unix(container.Created, 0)

		var portPrivate int
		var portPublic int
		var portString string
		for _, port := range container.Ports {
			portPrivate = int(port.PrivatePort)
			portPublic = int(port.PublicPort)
			portString = strconv.Itoa(portPrivate) + ":" + strconv.Itoa(portPublic)
		}

		var row = TableRow{container.Image, createdAt.Format(time.RFC822), portString}
		rows = append(rows, row)
	}

	renderTemplate(buf, "home.html", rows)

	return buf.String()
}
