package smith

import (
	"testing"
	"bytes"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestTemplatingCanParseDataToTemplate(t *testing.T) {
	loadTemplateConfig()
	loadTemplates()

	container := new(Container)
	createdAt, _ := time.Parse(time.RFC3339, "2018-08-08T22:00:00.0+02:00")
	container.ID = "fake-uuid"
	container.Name = "container_name"
	container.Ip = "some_ip"
	container.CreatedAt = createdAt
	container.Ports = "8080:8080"
	container.Status = true

	// Anonymous struct
	data := struct {
		Container *Container
	}{
		container,
	}
	buf := bytes.NewBufferString("")
	renderTemplate(buf, "container.html", data)

	assert.Contains(t, buf.String(), `action="/container/fake-uuid"`)
}
