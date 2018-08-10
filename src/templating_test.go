package smith

import (
	"testing"
	"bufio"
	"bytes"
)

func TestTemplatingCanParseDataToTemplate(t *testing.T) {
	var buffer bytes.Buffer
	writer := bufio.NewWriter(&buffer)
	loadTemplateConfig()
	loadTemplates()
	data := map[string]string{
		"Name": "moriorgames",
	}
	renderTemplate(writer, "sample.html", data)

	t.Logf("Result buffer: %s", buffer.String())

	//assert.Equal(t, "my name is moriorgames", b.String())
}
