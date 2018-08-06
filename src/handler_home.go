package smith

import (
	"fmt"
	"bytes"
	"net/http"
)

func renderHomeContent() string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)
	renderTemplate(buf, "home.html", nil)

	return buf.String()
}

func ViewHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, renderHomeContent())
}
