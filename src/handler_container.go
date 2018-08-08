package smith

import (
	"fmt"
	"bytes"
	"net/http"
	"github.com/gorilla/mux"
)

func renderContainerContent(id string) string {
	loadTemplateConfig()
	loadTemplates()
	buf := new(bytes.Buffer)
	renderTemplate(buf, "container.html", nil)

	return buf.String()
}

func ViewContainer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	fmt.Fprintf(w, renderContainerContent(id))
}

func UpdateContainer(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	router := CreateRouter()
	newUrl, _ := router.Get("app_edit_container").URL("id", id)

	// Save container

	w.WriteHeader(http.StatusSeeOther)
	http.Redirect(w, r, newUrl.String(), http.StatusSeeOther)
}
