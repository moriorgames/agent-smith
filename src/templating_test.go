package smith

import (
	"testing"
		"bytes"
		"github.com/stretchr/testify/assert"
)

func TestTemplatingCanParseDataToTemplate(t *testing.T) {
	loadTemplateConfig()
	loadTemplates()
	data := map[string]string{
		"ContainerID": "123456",
	}
	buf := bytes.NewBufferString("")
	renderTemplate(buf, "container.html", data)

	assert.Contains(t, buf.String(), `action="/container/123456"`)
}
