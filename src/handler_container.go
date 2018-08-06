package smith

import (
	"fmt"
	"bytes"
	"net/http"
)

func renderContainerContent() string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)
	renderTemplate(buf, "container.html", nil)

	return buf.String()
}

func ViewContainer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, renderContainerContent())
}
